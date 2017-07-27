-- noinspection SqlNoDataSourceInspectionForFile

-- table order convention - new entities as late as possible,
-- tables with FOREIGN KEY as earlier as possible 
CREATE TYPE status AS ENUM ('opened', 'started', 'ended');

CREATE TABLE project (
    id                SERIAL                                ,
    created           TIMESTAMP WITH TIME ZONE      NOT NULL,
    description       TEXT                          NOT NULL,
    logo              TEXT                          NOT NULL,
    name              TEXT                          NOT NULL,
    status            status                        NOT NULL,
    tags              TEXT[]           NOT NULL DEFAULT '{}',
    githuburl         TEXT                          NOT NULL,

    CONSTRAINT project_pk PRIMARY KEY (id)
);

CREATE TABLE "user" (
    id                SERIAL                                ,
    nickname          TEXT                          NOT NULL,
    description       TEXT                          NOT NULL,
    avatar            TEXT                          NOT NULL,

    CONSTRAINT user_pk PRIMARY KEY (id)
);

CREATE TABLE project_user (
    id                SERIAL                                ,
    project_id        INT                           NOT NULL,
    user_id           INT                           NOT NULL,
    signed_date       TIMESTAMP WITH TIME ZONE      NOT NULL,
    progress          INT                           NOT NULL,

    CONSTRAINT project_user_pk PRIMARY KEY (id),
    CONSTRAINT project_user_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
    CONSTRAINT project_user_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id),
    CONSTRAINT project_user_uq UNIQUE (project_id, user_id)
);

CREATE TABLE project_enroll (
    id                SERIAL                                ,
    project_id        INT                           NOT NULL,
    user_id           INT                           NOT NULL,
    enrolling_message TEXT                          NOT NULL,
    time              TIMESTAMP WITH TIME ZONE      NOT NULL,

    CONSTRAINT project_enroll_pk PRIMARY KEY (id),
    CONSTRAINT project_enroll_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
    CONSTRAINT project_enroll_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id),
    CONSTRAINT project_enroll_uq UNIQUE (project_id, user_id)
);

-- type is one of email, tel, vk, icq
CREATE TABLE contact_type (
    id                SERIAL                                ,
    type              TEXT                          NOT NULL,
    
    CONSTRAINT contact_type_pk PRIMARY KEY (id)
);

CREATE TABLE user_contact (
    id                SERIAL                                ,
    contact           TEXT                          NOT NULL,
    type_id           INT                           NOT NULL,
    user_id           INT                           NOT NULL,
    
    CONSTRAINT user_contact_pk PRIMARY KEY (id),
    CONSTRAINT user_contact_fk0 FOREIGN KEY (type_id) REFERENCES contact_type(id),
    CONSTRAINT user_contact_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
);

-- project tasks
CREATE TABLE task (
    id                SERIAL                                ,
    title             TEXT                          NOT NULL,
    description       TEXT                          NOT NULL,
    numberOfTask      INT                           NOT NULL,
    tags              TEXT[]           NOT NULL DEFAULT '{}',
    priority          INT                           NOT NULL,
    project_id        INT                           NOT NULL,
    project_user_id   INT                           NOT NULL,
    
    CONSTRAINT task_pk PRIMARY KEY (id),
    CONSTRAINT task_fk0 FOREIGN KEY (project_id) REFERENCES project(id),
    CONSTRAINT task_fk1 FOREIGN KEY (project_user_id) REFERENCES project_user(id)
);

CREATE TABLE comment (
    id                SERIAL                                ,
    text              TEXT                          NOT NULL,

    CONSTRAINT comment_pk PRIMARY KEY (id)
);

CREATE TABLE user_comment (
    id                SERIAL                                ,
    -- в случае цитирования или комментирования комментария
    comment_id        INT                           NOT NULL,
    date              TIMESTAMP WITH TIME ZONE      NOT NULL,
    user_id           INT                           NOT NULL,

    CONSTRAINT user_comment_pk PRIMARY KEY (id),
    CONSTRAINT user_comment_fk0 FOREIGN KEY (comment_id) REFERENCES comment(id),
    CONSTRAINT user_comment_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
);

CREATE TABLE course (
    id                SERIAL                                ,
    description       TEXT                          NOT NULL,
    logo              TEXT                          NOT NULL,
    rating            REAL                          NOT NULL,
    title             TEXT                          NOT NULL,

    CONSTRAINT course_pk PRIMARY KEY (id)
);

CREATE TABLE user_course (
    id                SERIAL                                ,
    course_id         INT                           NOT NULL,
    date              TIMESTAMP WITH TIME ZONE      NOT NULL,
    progress          INT                           NOT NULL,
    user_id           INT                           NOT NULL,
    
    CONSTRAINT user_course_pk PRIMARY KEY (id),
    CONSTRAINT user_course_fk0 FOREIGN KEY (course_id) REFERENCES course(id),
    CONSTRAINT user_course_fk1 FOREIGN KEY (user_id) REFERENCES "user"(id)
);

-- course statistics
CREATE TABLE statistics (
    id                SERIAL                                ,
    course_id         INT                           NOT NULL,
    hours             INT                           NOT NULL,

    CONSTRAINT statistics_pk PRIMARY KEY (id),
    CONSTRAINT statistics_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
);

-- recommended course for other course
CREATE TABLE recommended_course (
    id                SERIAL                                ,
    course_id         INT                           NOT NULL,
    link              TEXT                          NOT NULL,

    CONSTRAINT recommended_course_pk PRIMARY KEY (id),
    CONSTRAINT recommended_course_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
);

-- course lesson
CREATE TABLE lesson (
    id                SERIAL                                ,
    course_id         INT                           NOT NULL,
    description       TEXT                          NOT NULL,
    rating            INT                           NOT NULL,
    title             TEXT                          NOT NULL,

    CONSTRAINT lesson_pk PRIMARY KEY (id),
    CONSTRAINT lesson_fk0 FOREIGN KEY (course_id) REFERENCES course(id)
);

-- lesson practise
CREATE TABLE practise (
    id                SERIAL                                ,
    description       TEXT                          NOT NULL,
    lesson_id         INT                           NOT NULL,
    
    CONSTRAINT practise_pk PRIMARY KEY (id),
    CONSTRAINT practise_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
);

-- lesson video
CREATE TABLE video (
    id                SERIAL                                ,
    lesson_id         INT                           NOT NULL,
    link              TEXT                          NOT NULL,

    CONSTRAINT video_pk PRIMARY KEY (id),
    CONSTRAINT video_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
);

-- lesson task
CREATE TABLE test (
    id                SERIAL                                ,
    lesson_id         INT                           NOT NULL,
    title             TEXT                          NOT NULL,

    CONSTRAINT test_pk PRIMARY KEY (id),
    CONSTRAINT test_fk0 FOREIGN KEY (lesson_id) REFERENCES lesson(id)
);

CREATE TABLE tests_task (
    id                SERIAL                                ,
    question          TEXT                          NOT NULL,
    test_id           INT                           NOT NULL,

    CONSTRAINT tests_task_pk PRIMARY KEY (id),
    CONSTRAINT tests_task_fk0 FOREIGN KEY (test_id) REFERENCES test(id)
);

-- test question
CREATE TABLE variant (
    id                SERIAL                                ,
    correct_answer    BOOL                          NOT NULL,
    tests_task_id     INT                           NOT NULL,
    text              TEXT                          NOT NULL,

    CONSTRAINT variant_pk PRIMARY KEY (id),
    CONSTRAINT variant_fk0 FOREIGN KEY (tests_task_id) REFERENCES tests_task(id)
);

CREATE TABLE news (
    id                SERIAL                                ,
    created           TIMESTAMP WITH TIME ZONE      NOT NULL,
    description       TEXT                          NOT NULL,
    image             TEXT                          NOT NULL,
    edited            TIMESTAMP WITH TIME ZONE      NOT NULL,
    tags              TEXT[]           NOT NULL DEFAULT '{}',
    title             TEXT                          NOT NULL,

    CONSTRAINT news_pk PRIMARY KEY (id)
);
