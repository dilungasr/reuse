# reuse

**reuse** provides a simple way to automate changes to be made on your current project folder so you can reuse it for the next project.

- Perform project-wise replacement to any text in your code
- Automate your project setup commands
- Delete files you don't want in the next project
- Specify extensions to modify contents on
- Ignore files you don't want to modify
- Use terminal interactive mode or define everying in config file
- Supports yaml, toml and json... There you have it!

Don't start from scratch, don't get stressed...
Just **reuse** it :)

> ## Installing

If you're using golang and you have it on your machine, then just run,

```shell
go install github.com/dilungasr/reuse
```

If golang is not your favor, then I have got portable binaries for you!
Just click <a href="https://github.com/dilungasr/reuse/raw/master/bin/reuse">here</a> to only download the binaries ... Or clone the entire branch and get your binaries inside the `/bin` folder...Ooohww come one! that is an overkill, right? The choice is yours!

After having your binaries on your machine, the next step should be adding them to your `$PATH` so you can be able to run them from anywhere on the terminal.

> ## Intro - Running

There are two modes to run **reuse** which are **interactive** and **config** mode.

> ## Interactive Mode

With **interactive mode**, you are able to define your configuration directly on the terminal via a series interactive prompts.
To run **reuse** in interactive mode just `cd` to the project folder and run:

```shell
reuse -i
# OR
reuse --interact
```

Now you will be prompted to provide your configuration with a series of straight forward prompts.
Here are prompts and what they mean,

`Extensions (use spaces to separate):`

- Enter extensions (type) of files to modify the contents on. For example `go js py`

---

`Files/folders to ignore (use spaces to separate):`

- Specify files or folders to ignore. For example `main.go user.go person.go`

---

```shell
   Define your replacements (in old: new pair per line)
   Enter q to finish:
```

- Enter your replacements in `old: new` syntax line by line. Use q and `Enter` to finish defining your placements. For example
  ```shell
  github.com/dilungasr/goreuse: github.com/dilungasr/reuse
  Saul: Paul
  Abraham: Ibrahim
  q #you're done!
  ```

---

`Files/folders to delete (use spaces to separate): `

- List of files or folders to delete. For example `main.go folder1`

---

```shell
   Commands to run (sequentially - line by line)
   Enter q to finish:
```

- Provide a list of setup commands to run at the end after modifications. List them one command per line and they will be run sequentially. Use q and `Enter` to finish listing. For example
  ```shell
  git init
  git remote add https://github.com/dilungasr/reuse
  q #you're done!
  ```

---

> ## Config Mode

As mentioned before, the second mode in **config mode**. With config mode you are able to provide your configuration in your config file and **reuse** will read your options and use them in automating the changes.

**reuse** currently supports json, yaml and toml config files.

#### Defining your configuration.

The options are exactly similar to those your were prompted to provide in **iteractive mode**. The little difference is that now you define them in the config file which is a greet way of automating stuffs.

To start the game just choose your flavor... and create **reuse.json** or **reuse.yaml** or **reuse.yml** or **reuse.toml** depending on your choice at the root of the project folder.

Here are examples of using config files to define your configuration.

##### reuse.toml

```toml

# files to delete
del = ["go.sum", "go.mod"]
# extensions to modify contents on
ext = ["go"]
# files to ignore in the iterations
ignore = ["reuse.toml"]
# commnds to run at the end
run = [
  "go mod init 'github.com/dilungasr/reuse'",
  "go mod tidy -v",
  "echo 'Happy coding!'",
]
# replacements in old = new pair
[rep]
Text1 = "NewText1"
Text2 = "NewText2"
"github.com/dilungasr/goreuse" = "github.com/dilungasr/reuse"

```

##### reuse.yaml or reuse.yml

```yml
# replacements in old: new pair
rep:
  Text1: NewText1
  Text2: NewText2
  github.com/dilungasr/goreuse: "github.com/dilungasr/reuse"
# files to delete
del:
  - go.sum
  - go.mod
  - .git
# commnds to run at the end
run:
  - go mod init "github.com/dilungasr/reuse"
  - go mod tidy -v
  - echo "Happy coding!"
# extensions to modify contents on
ext:
  - go
# files to ignore in the iterations
ignore:
  - reuse.yml
```

##### reuse.json

```json
{
  "rep": {
    "Text1": "NewText1",
    "Text2": "NewText2",
    "github.com/dilungasr/goreuse": "github.com/dilungasr/reuse"
  },
  "del": ["go.sum", " go.mod", ".git"],
  "run": [
    "go mod init 'github.com/dilungasr/reuse'",
    "go mod tidy -v",
    "echo 'Happy coding!'"
  ],
  "ext": ["go"],
  "ignore": ["reuse.json"]
}
```

#### Running in config mode

At the simplest level, all you have to do is,

```shell
reuse
```

#### Using other names than "reuse" or location than project's root.

Luckily, **reuse** is extremely unopinionated... You can easily point to your supported config file by using the `-u` flag as simple as:

```shell
reuse -u ~/projects/myproject/myconfig.yaml
#OR
reuse --use ~/projects/myproject/myconfig.yaml
```

or even easier, with the help of **@** you can refer to the root of your project. For example:

```shell
reuse -u @myconfig.toml #reads myconfig.toml in the root of our project folder
#OR
reuse -u @subdir/myconfig.toml #reads myconfig.toml in the subdir folder(subdirectory)
```

> ## Running Reuse Outside the Project Folder

Oooh yeaah! it is not a must to run **reuse** inside the project folder! You can easily run it from anywhere and pass the path to your target project folder.

To do so, just provide the project path as the argument to the **reuse** command,

```shell
reuse ~/projects/myprojects/project1
```

Everything else is just similar to running **reuse** inside the project folder. In fact, running **reuse** inside the project folder is exactly similar to:

```shell
reuse .
```

But in this case you don't have to explicitly specify the path(.) since **reuse** does this for you!

> ## Configuration Precendece

Well, understanding how **reuse** picks configuration mode or how it works on finding the config files is crucial to avoid confusion.

Here is a flow...

- **reuse** first checks if mode is interactive...then runs in interactive mode
- If not interactive, then it checks if config file path is provided via **-u** flag. If so, it reads from it and fails the program if the file not found
- If config file not provided, **reuse** will attempt to find for **reuse file with supported config extension** in the root directory.
- If an attempt fails (no config file found), **reuse** will decide to run in **interactive mode**.

The precedence level ranking:

1. Interactive
2. explicitly provided config file via **-u**
3. reuse.js or reuse.yaml or reuse.yaml or reuse.toml

# Contributing

Everyone with Golang skills is warmly welcomed!
