#!/bin/bash
#
# Some backend tasks
# See $helpmsg for further info

export STUDIT
STUDIT=$(pwd)

if [ -z "${GOPATH}" ]; then
    export GOPATH="${GOPATH}:${STUDIT}"
else
    export GOPATH="${STUDIT}"
fi


readonly helpmsg="I'm script for running microservices, updating dependencies and databases

Usage: ./run.sh [command]

Commands:
    run (default)  runs specified services (or all by default)
    resetdb        resets specified databases (or all by default)
    install        installs or updates microservices' dependencies
    help           prints this help message

'run' tries to run in tmux first, then in some popular terminal emulators
If there's nor tmux neither any known terminal emulator recognized, runs
all microservices in the same shell

When running in tmux, following functions are exported:
    resetdb      same as ./run.sh resetdb
    sms_install  same as ./run.sh install
    sms_start    (re)start specified services  (or all by default)
    sms_stop     stop specified services  (or all by default)
"

readonly dependencies=("go" "bee" "dep")

installed() {
    hash $1 2>/dev/null
}

_install_deps() {
    local services=("main-service" "auth-service" "data-service" "file-service")

    echo '\e[93mUpdating dependencies. It may take \e[3mreally\e[0m \e[93mlong
(see https://github.com/golang/dep/blob/master/docs/FAQ.md#why-is-dep-slow) \e[90m'
    cd "${STUDIT}" || return
    for service in "${services[@]}"; do
        cd "${service}" || return
        echo "\e[93mRunning 'dep ensure -v' for ${service}:\e[0m"
        dep ensure -v || echo "\e[31m'dep ensure' exited with non-zero status
see above output for details\e[0m"
        cd "${STUDIT}" || return
    done
}

# Re-create databases
_resetdb() {
    cd "${STUDIT}" || return

    # Parse arguments, fallback to all services
    local args=()
    for arg in "$@"; do
        args+=(${arg})
    done

    if [ $# -eq 0 ]; then
        args=('auth' 'file' 'data')
    fi

    # Re-create databases for specified arguments
    for arg in "${args[@]}"; do
        case ${arg} in
            auth|auth-service)
                cd "auth-service/schema" || return
                ./init.sh
                cd "${STUDIT}" || return
                ;;
            data|data-service)
                cd "data-service/schema" || return
                ./init.sh
                cd "${STUDIT}" || return
                ;;
            file|file-service)
                cd "file-service/schema" || return
                ./init.sh
                cd "${STUDIT}" || return
                ;;
            *)
                echo "Wrong argument: ${arg}"
                ;;
        esac
    done
}

start_tmux() {
    # unique session name
    SESSION="studit-backend[$$]"

    tmux start-server
    tmux -2 new-session -d -s "${SESSION}"

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
    tmux select-pane -t 1
    data_id=$(tmux display-message -p "#{pane_id}")
    tmux select-pane -t 2
    auth_id=$(tmux display-message -p "#{pane_id}")
    tmux select-pane -t 3
    file_id=$(tmux display-message -p "#{pane_id}")
    tmux select-pane -t 4
    console_id=$(tmux display-message -p "#{pane_id}")

    # Export functions to bottom pane
    # Usage: tmux_export_function fname
    tmux_export_function() {
        for func in "$@"; do
            tmux send-keys -t "${console_id}" " $(declare -f ${func})"
            tmux send-keys -t "${console_id}" C-m
        done
        tmux send-keys -t "${console_id}" " clear"
        tmux send-keys -t "${console_id}" C-m
    }

    # Clear pane
    # Usage: tmux_clear pane_id
    tmux_clear() {
        tmux send-keys -t "$1" " clear"
        tmux send-keys -t "$1" C-m
    }

    # Exec command in pane
    # Usage: tmux_exec pane_id command
    tmux_exec() {
        tmux send-keys -t "$1" " ${*:2}"
        tmux send-keys -t "$1"  C-m
    }

    # Export variable to bottom pane
    # Usage: tmux_export_variable variable value
    tmux_export_variable() {
        tmux send-keys -t "${console_id}" "  $1=$2"
        tmux send-keys -t "${console_id}" C-m
        tmux_clear "${console_id}"
    }

    # Export pane id's for running services to stop/start them via in-tmux functions
    tmux_export_variable file_id "${file_id}"
    tmux_export_variable main_id "${main_id}"
    tmux_export_variable auth_id "${auth_id}"
    tmux_export_variable data_id "${data_id}"

    tmux_stop() {
        for pane in "$@"; do
            tmux respawn-pane -k -t "${pane}" -c "${STUDIT}"
        done
    }

    kill_tmux() {
        tmux kill-session -t "${SESSION}"
    }


    run_main() {
        cd ${STUDIT} || return
        tmux_exec "${main_id}" "bee run main-service"
    }

    run_data() {
        cd ${STUDIT} || return
        tmux_exec "${data_id}" "bee run -downdoc=true -gendoc=true data-service"
    }

    run_auth() {
        cd ${STUDIT} || return
        tmux_exec "${auth_id}" "bee run -downdoc=true -gendoc=true auth-service"
    }

    run_file() {
        cd ${STUDIT} || return
        tmux_exec "${file_id}" "bee run -downdoc=true -gendoc=true file-service"
    }

    run_services() {
        run_auth
        run_file
        run_data
        run_main
    }


    sms_stop() {
        local args=()
        for arg in "$@"; do
            args+=("${arg}")
        done

        if [ $# -eq 0 ]; then
            args=('auth' 'main' 'file' 'data')
        fi

        for arg in "${args[@]}"; do
            case ${arg} in
                auth|auth-service)
                    tmux_stop ${auth_id};;
                data|data-service)
                    tmux_stop ${data_id};;
                file|file-service)
                    tmux_stop ${file_id};;
                main|main-service)
                    tmux_stop ${main_id};;
                *)
                    echo "No such service: ${arg}";;
            esac
        done
    }

    sms_start() {
        local args=()
        for arg in "$@"; do
            args+=(${arg})
        done

        if [ $# -eq 0 ]; then
            args=('auth' 'main' 'file' 'data')
        fi

        for arg in "${args[@]}"; do
            case ${arg} in
                auth|auth-service)
                    run_auth;;
                data|data-service)
                    run_data;;
                file|file-service)
                    run_file;;
                main|main-service)
                    run_main;;
                *)
                    echo "No such service: ${arg}";;
            esac
        done
    }

    # Run init scripts for services' databases
    # Usage: resetdb service1 service2 ...
    resetdb() {
        sms_stop "$@"
        _resetdb "$@"
        sms_start "$@"
    }

    sms_install() {
        sms_stop
        _install_deps
        sms_start
    }

    # Export functions to tmux
    tmux_export_function run_main run_data run_file run_auth run_services tmux_clear tmux_exec tmux_stop kill_tmux _resetdb resetdb sms_start sms_stop sms_install _install_deps

    # Kill session on Ctrl-D
    tmux_exec ${console_id} trap kill_tmux EXIT
    tmux_clear

    run_services

    tmux -2 attach-session -t ${SESSION}
}

interactive() {
    if installed tmux; then
        start_tmux
    elif installed xfce4-terminal; then
        xfce4-terminal --title="StudiIT data-service" --command="${STUDIT}/data-service/run-data-service.sh" &
        xfce4-terminal --title="StudiIT auth-service" --command="${STUDIT}/auth-service/run-auth-service.sh" &
        xfce4-terminal --title="StudiIT file-service" --command="${STUDIT}/file-service/run-file-service.sh" &
        xfce4-terminal --title="StudiIT main-service" --command="${STUDIT}/main-service/run-main-service.sh"
    elif installed konsole; then
        konsole -e "${STUDIT}/data-service/run-data-service.sh" &
        konsole -e "${STUDIT}/auth-service/run-auth-service.sh" &
        konsole -e "${STUDIT}/file-service/run-file-service.sh" &
        konsole -e "${STUDIT}/main-service/run-main-service.sh"
    elif installed gnome-terminal; then
        gnome-terminal -- "${STUDIT}/data-service/run-data-service.sh" &
        gnome-terminal -- "${STUDIT}/auth-service/run-auth-service.sh" &
        gnome-terminal -- "${STUDIT}/file-service/run-file-service.sh" &
        gnome-terminal -- "${STUDIT}/main-service/run-main-service.sh"
    else
        echo "You don't have any known terminal emulator or multiplexer installed."
        echo "Running services in one shell"
        main-service/run-main-service.sh &
        auth-service/run-auth-service.sh &
        file-service/run-file-service.sh &
        data-service/run-data-service.sh
    fi
}

# Check installed dependencies
for dep in "${dependencies[@]}"; do
    if ! installed ${dep}; then
        echo "You don't have ${dep} in yout PATH. Exiting"
        exit 1
    fi
done

if [ $# -eq 0 ]; then
    job="run"
else
    job=$1
fi

case ${job} in
    run)
        interactive
        ;;
    resetdb)
        _resetdb "${*:2}"
        ;;
    i|install)
        _install_deps "${*:2}"
        ;;
    *)
        echo "${helpmsg}"
        ;;
esac
