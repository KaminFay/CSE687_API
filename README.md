# CSE687_API
### API For Final Project of CSE687 at Syracuse University

# Description
- This API is written to be the middle man between a C# GUI and C++ Backend for a DLL test harness. Currently the API is hosted at www.kaminfay.com
- The API connects to a postgres database on the server to store the data that needs to be passed back and forth

# Endpoints

- http://www.kaminfay.com/cse687/sendFunctions
  - Input Body -- JSON -- Example:
    - ```
      {
        "FuncName":"Run_Test",
        "DllName":"dll_long_delay.dll",
        "DllPath":"C:\\Users\\Kamin\\Documents\\GitHub\\CSE687\\Testing\\dll_files\\dll_long_delay.dll"
      }
      ```
  - Output -- No Data Returned

- http://www.kaminfay.com/cse687/recieveFunctions
  - Input Body -- No Data Sent
  - Output -- JSON -- Example:
    - ```
      [
      {
          "ID": 261,
          "FuncName": "Run_Test",
          "DllName": "dll_long_delay.dll",
          "DllPath": "C:\\Users\\Kamin\\Documents\\GitHub\\CSE687\\Testing\\dll_files\\dll_long_delay.dll"
      },
      {
          "ID": 262,
          "FuncName": "Run_Test",
          "DllName": "dll_long_delay.dll",
          "DllPath": "C:\\Users\\Kamin\\Documents\\GitHub\\CSE687\\Testing\\dll_files\\dll_long_delay.dll"
      }
      ```
      
- http://www.kaminfay.com/cse687/sendResults
  - Input Body -- JSON -- Example:
    - ```
      {
      "DllName": "C:\\Users\\Kamin\\Documents\\GitHub\\CSE687\\Testing\\dll_files\\dll_long_delay.dll",
      "DllPath": "C:\\Users\\Kamin\\Documents\\GitHub\\CSE687\\Testing\\dll_files\\dll_long_delay.dll",
      "EndTime": "2020-06-14 03:29:14",
      "Exception": "",
      "FuncName": "Run_Test",
      "ID": 185,
      "Result": false,
      "StartTime": "2020-06-14 03:29:09"
      }
      ```
  - Output -- No Data Returned
  
- http://www.kaminfay.com/cse687/sendFunctions
  - Input Body -- JSON -- Example:
    - ```
      {
      "ID": 185
      }
      ```
  - Output -- JSON -- Example:
    - ```
      {
      "DllName": "C:\\Users\\Kamin\\Documents\\Garbage\\Dll_Tester\\auto_pass_delayed.dll",
      "DllPath": "C:\\Users\\Kamin\\Documents\\Garbage\\Dll_Tester\\auto_pass_delayed.dll",
      "EndTime": "2020-06-15 17:38:29",
      "Exception": "",
      "FuncName": "auto_pass_delayed",
      "ID": 267,
      "Result": true,
      "StartTime": "2020-06-15 17:38:27"
      }
      ```
      
# Database Structure:

```
 Schema |         Name          |   Type   |  Owner
--------+-----------------------+----------+----------
 public | finished_tests        | table    | postgres
 public | test_functions        | table    | postgres
```

# Table Structure:

```
                                       Table "public.test_functions"
    Column     |          Type          | Collation | Nullable |                  Default
---------------+------------------------+-----------+----------+--------------------------------------------
 id            | integer                |           | not null | nextval('test_functions_id_seq'::regclass)
 function_name | character varying(100) |           | not null |
 dll_name      | character varying(100) |           |          |
 dll_path      | character varying(100) |           | not null |
Indexes:
    "test_functions_pkey" PRIMARY KEY, btree (id)
```

```
                      Table "public.finished_tests"
    Column     |          Type          | Collation | Nullable | Default
---------------+------------------------+-----------+----------+---------
 id            | integer                |           | not null |
 function_name | character varying(100) |           | not null |
 dll_name      | character varying(100) |           |          |
 dll_path      | character varying(100) |           | not null |
 bool_result   | boolean                |           | not null |
 exception     | character varying(255) |           | not null |
 start_time    | character varying(100) |           | not null |
 end_time      | character varying(100) |           | not null |
Indexes:
    "finished_tests_pkey" PRIMARY KEY, btree (id)
```
