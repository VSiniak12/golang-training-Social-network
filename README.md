# golang-training-Social-network
My theme is "Social network". 

## Task 3

### First part
1. Create your `PostgresSQL`database consisting of 5 tables and 2 lookup tables (without links to parent tables), if necessary, you can create third. Your database should have a third normal form.
Create a data schema in `draw.io` (you can in access), add the image with your schema to repo with task implementation.
2. Add SQL script that creates tables, and add 5 records to each table. Save the script to the `script.sql` file.
3. Required type of fields in the database:
- `number`
- `date`
- `string`
- `boolean`

### Second part
Write a Golang application for manipulating data from the database created above. Use Standard library or `Gorm` library.
What should be present in the code:
 - Create the necessary entities at your discretion. 
 - At least one JOIN (all other requests at your discretion);
 - Organize CRUD (create, read, update, delete) operations

### Third part
Add unit tests, code coverage must be 50% and higher.

### Fourth part (Optional)
Add migrations using goose library: https://github.com/pressly/goose


#### Let me know if you have any questions ;)
