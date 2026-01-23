UPDATE mdl_task_adhoc 
SET attemptsavailable = 1, 
    timestarted = NULL, 
    faildelay = 0, 
    nextruntime = 1 
WHERE classname LIKE '%send_user_notifications%';


#this is for moodle notifiacation restart


mysql -u sharda -p"sharda@2026@latest" sharda_db


# this is for moodle notification restart for jain


kubectl exec -n jo-moodle moodle-db-5b97798f77-pctmx -- mysql -u moodleuser -pmoodlepass moodledb -e "UPDATE acu_task_adhoc SET timestarted = NULL, nextruntime = UNIX_TIMESTAMP(), faildelay = 0 WHERE timestarted IS NOT NULL;"
 
 #this is for que in in quickmail

SELECT id, subject, sent_at, to_send_at, is_draft, is_sending, timecreated FROM acu_block_quickmail_messages ORDER BY id DESC LIMIT 5;


UPDATE acu_block_quickmail_messages SET is_sending = 0 WHERE id = 117;


php admin/cli/adhoc_task.php --id=197802

# this is the cmd where i can directly restore the sql dump into the pod
kubectl exec -i <pod-name> -- mysql -u <username> -p<password> <database-name> < ./path/to/your/dump.sql


# To update the "queue_quiz_open_notification_tasks_for_users" in moodle

UPDATE acu_task_adhoc 
SET nextruntime = UNIX_TIMESTAMP(), 
    faildelay = 0, 
    timestarted = NULL, 
    attemptsavailable = 50 
WHERE classname LIKE '%queue_quiz_open_notification_tasks_for_users%';


# this is the cmd to delete the task which are in failed

DELETE FROM acu_task_adhoc WHERE component LIKE '%quiz%' AND faildelay > 0;