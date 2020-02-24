# Knowledge base structure

Knowledge base (KB) is tree of records. Each record stored in separate directory.
Name of directory is record's short id. Root record in KB has empty short id.
Path to record relative to root record is full id.
Each record can contain other records within it. 
Level of nesting restricted only by file system.
Record is defined by its id and file `info.json` with meta information.
  
Example of file `info.json`

```json
{
    "create_date": "2020-02-06T17:38:17.242183+03:00",
    "title": "This is root of knowledge base",
    "description": ""
}
```

Description of `info.json` structure:
 - `create_date` contains date when record was created.
 - `title` very short description of the record, in few word.
 - `description` description of record in few sentences.