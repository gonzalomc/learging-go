## typed
## fast compile times
## garbage collected
## built-in concurrency
## compile to sandalone binaries



<!-- Sql Tables -->
create table if not exists surveys(
    id INT GENERATED ALWAYS AS IDENTITY,
    name varchar(200),
    status boolean default true,
    PRIMARY KEY(id)
);


create table if not exists question(
    id INT GENERATED ALWAYS AS IDENTITY,
    title varchar(200),
    survey_id int,
    PRIMARY KEY(id),
    CONSTRAINT fk_survey
        FOREIGN KEY(survey_id)
            REFERENCES surveys(id)
);