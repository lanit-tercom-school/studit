-- queries.sql

-- name: GetAllEnrolledOnProject
SELECT pe.enrolling_message,
        p.id AS project_id,
        p.name AS project_name,
         p.tags AS project_tags,
        p.status AS project_status,
        p.logo AS project_logo,
        p.description AS project_description,
         p.date_of_creation AS project_date,
        u.id AS user_id,
        u.nickname AS user_name,
         u.avatar AS user_avatar,
         u.description AS user_description
FROM public.project_user pu
INNER JOIN public.project_enroll pe
    ON pu.project_id=pe.project_id
INNER JOIN public.user u
    ON pe.user_id=u.id
INNER JOIN public.project p
    ON p.id=pe.project_id
WHERE pu.user_id=?