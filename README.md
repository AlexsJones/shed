# shed

```
                hhhhhhh                                             d::::::d
                h:::::h                                             d::::::d
                h:::::h                                             d::::::d
                h:::::h                                             d:::::d
    ssssssssss   h::::h hhhhh           eeeeeeeeeeee        ddddddddd:::::d
  ss::::::::::s  h::::hh:::::hhh      ee::::::::::::ee    dd::::::::::::::d
ss:::::::::::::s h::::::::::::::hh   e::::::eeeee:::::ee d::::::::::::::::d
s::::::ssss:::::sh:::::::hhh::::::h e::::::e     e:::::ed:::::::ddddd:::::d
 s:::::s  ssssss h::::::h   h::::::he:::::::eeeee::::::ed::::::d    d:::::d
   s::::::s      h:::::h     h:::::he:::::::::::::::::e d:::::d     d:::::d
      s::::::s   h:::::h     h:::::he::::::eeeeeeeeeee  d:::::d     d:::::d
ssssss   s:::::s h:::::h     h:::::he:::::::e           d:::::d     d:::::d
s:::::ssss::::::sh:::::h     h:::::he::::::::e          d::::::ddddd::::::dd
s::::::::::::::s h:::::h     h:::::h e::::::::eeeeeeee   d:::::::::::::::::d
 s:::::::::::ss  h:::::h     h:::::h  ee:::::::::::::e    d:::::::::ddd::::d
  sssssssssss    hhhhhhh     hhhhhhh    eeeeeeeeeeeeee     ddddddddd   ddddd
  ```

  Shell script scheduling.

  Automate repetitive tasks and save them as a `Shedfile`

  Hooya.

  ```
  clear      clear the current stack
  exit       exit the program
  help       display help
  list       list execution order
  load       Loads a local ShedFile into a schedule
  logs       logs from an execution
  push       push a k8s config-map path
  retry      retry a certain action based on index
  run        Starts running k8s config-map paths
  save       Saves out a new ShedFile
  ```

  # Example

  I want to automate a simple workflow for kubernetes deployment...

```
# In my kubernetes project directory

>>> push kubectl config view
Pushing -> kubectl
>>> push ls
Pushing -> ls
>>> list
+------+------------------+-----------+----------+------------+
| STEP | RESOURCE LOCATOR | VALIDATED | EXECUTED | SUCCESSFUL |
+------+------------------+-----------+----------+------------+
|    1 | kubectl          | ✓         | ✗        | ?          |
|    2 | ls               | ✓         | ✗        | ?          |
+------+------------------+-----------+----------+------------+
>>> push kubectl create -f .
Pushing -> kubectl
>>> list
+------+------------------+-----------+----------+------------+
| STEP | RESOURCE LOCATOR | VALIDATED | EXECUTED | SUCCESSFUL |
+------+------------------+-----------+----------+------------+
|    1 | kubectl          | ✓         | ✗        | ?          |
|    2 | ls               | ✓         | ✗        | ?          |
|    3 | kubectl          | ✓         | ✗        | ?          |
+------+------------------+-----------+----------+------------+
>>> save
Created new Shedfile...
>>> run
```
