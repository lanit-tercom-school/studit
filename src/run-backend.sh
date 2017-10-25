#!/usr/bin/env bash
# Run backend
# Tries to run services in byobu with tmux backend or pure tmux
# If tmux is not found, tries some popular terminal emulators
# Falls back to running all services in current shell
# When running in tmux, additional functions are defined:
# resetdb - recreates specified databases or all databases if none are specified
# sms_stop - stops specified services or all services if none are specified
# sms_start - (re)starts specified services or all services if none are specified

export STUDIT=`pwd`

installed() {
    command -v $1
}

start_tmux() {
    SESSION="studit-backend[$$]"

    if installed byobu; then
        tmux="byobu-tmux"
    else
        tmux="tmux"
    fi

    tmux start-server
    tmux -2 new-session -d -s $SESSION

    # Split window into three even horizontal panes
    tmux split-window  -v
    tmux split-window  -v
    tmux select-layout even-vertical

    # Split two upper panes horizontally
    # +-----+
    # |0 | 1|
    # +-----+
    # |2 | 3|
    # +-----+
    # |  4  |
    # +-----+
    tmux select-pane -t 0
    tmux split-window -h
    tmux select-pane -t 2
    tmux split-window -h

    # Find panel ids
    tmux select-pane -t 0
    main_id=$(tmux display-message -p "#{pane_id}")
    tmux set -g @main_id "$main_id"
    tmux select-pane -t 1
    data_id=$(tmux display-message -p "#{pane_id}")
    tmux select-pane -t 2
    auth_id=$(tmux display-message -p "#{pane_id}")
    tmux select-pane -t 3
    file_id=$(tmux display-message -p "#{pane_id}")
    tmux select-pane -t 4
    console_id=$(tmux display-message -p "#{pane_id}")

    # Export function to bottom pane
    # Usage: tmux_export_function fname
    tmux_export_function() {
        for func in "$@"
        do
            tmux send-keys -t $console_id " `declare -f $func`"
            tmux send-keys -t $console_id C-m
        done
        tmux send-keys -t $console_id " clear"
        tmux send-keys -t $console_id C-m
    }

    # Clear pane
    # Usage: tmux_clear pane_id
    tmux_clear() {
        tmux send-keys -t $1 " clear"
        tmux send-keys -t $1 C-m
    }

    # Exec command in pane
    # Usage: tmux_exec pane_id command
    tmux_exec() {
        tmux send-keys -t $1 " ${*:2}"
        tmux send-keys -t $1  C-m
    }

    # Export variable to bottom pane
    # Usage: tmux_export_variable variable value
    tmux_export_variable() {
        tmux send-keys -t $console_id "  $1=$2"
        tmux send-keys -t $console_id C-m
        tmux_clear $console_id
    }

    # Export pane id's for running services to stop/start them via in-tmux functions
    tmux_export_variable file_id $file_id
    tmux_export_variable main_id $main_id
    tmux_export_variable auth_id $auth_id
    tmux_export_variable data_id $data_id


    tmux_stop() {
        for pane in "$@"
        do
            tmux respawn-pane -k -t $pane -c $STUDIT
        done
    }

    kill_tmux() {
        tmux kill-session -t "$SESSION"
    }


    run_main() {
        tmux_exec $main_id "main-service/run-main-service.sh"
    }

    run_data() {
        tmux_exec $data_id "data-service/run-data-service.sh"
    }

    run_auth() {
        tmux_exec $auth_id "auth-service/run-auth-service.sh"
    }

    run_file() {
        tmux_exec $file_id "file-service/run-file-service.sh"
    }

    run_services() {
        run_auth
        run_file
        run_data
        run_main
    }


    sms_stop() {
        cd $STUDIT

        args=("${@[@]}")

        if [ $# -eq 0 ]; then
            args+=('auth')
            args+=('file')
            args+=('data')
            args+=('main')
        fi

        for arg in $args
        do
            case $arg in
                auth|auth-service)
                    tmux_stop $auth_id;;
                data|data-service)
                    tmux_stop $data_id;;
                file|file-service)
                    tmux_stop $file_id;;
                main|main-service)
                    tmux_stop $main_id;;
                *)
                    echo "No such service: $arg";;
            esac
        done
    }

    sms_start() {
        cd $STUDIT

        sms_stop $@
        args=("${@[@]}")

        if [ $# -eq 0 ]; then
            args+=('auth')
            args+=('file')
            args+=('data')
            args+=('main')
        fi

        for arg in $args
        do
            case $arg in
                auth|auth-service)
                    run_auth;;
                data|data-service)
                    run_data;;
                file|file-service)
                    run_file;;
                main|main-service)
                    run_main;;
                *)
                    echo "No such service: $arg";;
            esac
        done
    }

    # Run init scripts for services' databases
    # Usage: resetdb service1 service2 ...
    resetdb() {
        cd $STUDIT

        args=("${@[@]}")

        if [ $# -eq 0 ]; then
            args+=('auth')
            args+=('file')
            args+=('data')
        fi

        for arg in $args
        do
            case $arg in
                auth|auth-service)
                    tmux_stop $auth_id
                    cd auth-service/schema
                    ./init.sh
                    cd $STUDIT
                    run_auth;;
                data|data-service)
                    tmux_stop $data_id
                    cd data-service/schema
                    ./init.sh
                    cd $STUDIT
                    run_data;;
                file|file-service)
                    tmux_stop $file_id
                    cd file-service/schema
                    ./init.sh
                    cd $STUDIT
                    run_file;;
                *)
                    echo "Wrong argument: $arg";;
            esac
        done
    }

    # Export functions to tmux
    tmux_export_function run_main run_data run_file run_auth run_services tmux_clear tmux_exec tmux_stop kill_tmux resetdb sms_start sms_stop

    # Kill session on Ctrl-D
    tmux_exec $console_id trap kill_tmux EXIT
    tmux_clear

    run_services

    tmux -2 attach-session -t $SESSION
}

if installed tmux; then
    start_tmux
elif installed xfce4-terminal; then
    xfce4-terminal --title="StudiIT data-service" --command="$STUDIT/data-service/run-data-service.sh" &
    xfce4-terminal --title="StudiIT auth-service" --command="$STUDIT/auth-service/run-auth-service.sh" &
    xfce4-terminal --title="StudiIT file-service" --command="$STUDIT/file-service/run-file-service.sh" &
    xfce4-terminal --title="StudiIT main-service" --command="$STUDIT/main-service/run-main-service.sh"
elif installed konsole; then
    konsole -e "$STUDIT/data-service/run-data-service.sh" &
    konsole -e "$STUDIT/auth-service/run-auth-service.sh" &
    konsole -e "$STUDIT/file-service/run-file-service.sh" &
    konsole -e "$STUDIT/main-service/run-main-service.sh"
elif installed gnome-terminal; then
    gnome-terminal -- "$STUDIT/data-service/run-data-service.sh" &
    gnome-terminal -- "$STUDIT/auth-service/run-auth-service.sh" &
    gnome-terminal -- "$STUDIT/file-service/run-file-service.sh" &
    gnome-terminal -- "$STUDIT/main-service/run-main-service.sh"
else
    echo "You don't have any known terminal emulator or multiplexer installed."
    echo "Running services in one shell"
    main-service/run-main-service.sh &
    auth-service/run-auth-service.sh &
    file-service/run-file-service.sh &
    data-service/run-data-service.sh
fi
