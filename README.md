![forgit logo](/forgit_md_logo.png)

# Forgit CLI
This works along with the Forgit Web App. This lives on you local machine and automates your git add, commit, pull, and push workflow.
***

### Commands

#### forgit init
You need to be connected to the internet for this command. Sets up computer environment. It will create a hidden file in $HOME directroy called .forgitConf.json and a Forgit directory from the path you type in.

```
$ forgit init
```
**Or**
```
$ forgit i
```  
* **Output**  
    <> Your Current Absolute Path is -> /Users/user/current/path/  
    <>  Path cannot contain Forgit name.  
    <> Enter Absolute path where you want the Forgit directory [ Enter For Here ]: /PATH/  
    <> Enter UUID from Forgit Online Terminal Page: YOUR ID

___

#### forgit start
This command has a few params you can pass in. In general it will start the app with the setting workspace that is already selected on the web interface.  
  * **General**  

    ```
    $ forgit start
    ```  

    * **Output**:  
        To select a Workspace  
        -->  forgit start group GROUP-NAME  
        This session will have the following settings:  
        Setting Name:  General  
        Commit Time:  1  
        Push Time:  2  

        ___

  *  **Offline or single session**  
    Commit time minutes param: **-c**  
    Push time minutes param: **-p**

    ```
    $ forgit start -c 5 -p 30
    ```

    * **Output**:  
        To select a Workspace  
        -->  forgit start group GROUP-NAME  
        This session will have the following settings:  
        Setting Name:  forgitDefault  
        Commit Time:  2  
        Push Time:  2

        ___

  * **Setting Workspace**  
    The workspace name has to be spelled exact or it will not work.  
    select workspace param: **g**  

    ```
    $ forgit start g General
    ```

    * **Output**:  
        This session will have the following settings:  
        Setting Name:  General  
        Commit Time:  1  
        Push Time:  2

___

#### forgit stop  

 * **Output**:   
  To stop the app you must do ONE of the following:  
  1. Close the forgit shell window.  
  2. Control-c in the forgit window.

___

#### forgit help
Help will show possible commands and how to use them.  

```
$ forgit help
```

**Or**

```
$ forgit h
```
