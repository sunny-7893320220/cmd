# zammad login in to psql 
psql -U zammad -d zammad_production 

password : zammad

# this is the query to get the count to ticket in zammad

SELECT ts.name AS state, COUNT(t.id) AS count FROM tickets t JOIN ticket_states ts ON t.state_id = ts.id GROUP BY ts.name;


sudo gmode-toggle