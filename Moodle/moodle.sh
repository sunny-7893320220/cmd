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


php admin/cli/adhoc_task.php --id=40031
