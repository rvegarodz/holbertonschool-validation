# Awesome Inc. Website
This repository contains the source code for the Awesome Inc. website and the instructions on how to use it.

## Prerequisites
To use the resources provided, you need to have the following tools installed in your system:

    - GoHugo version 0.84.0
    - GNU Make version 3.81 or 4.0

## Lifecycle
The development lifecycle follows three main steps:

    - Build: Generate the website from the markdown and configuration files in the directory named dist.
    - Clean: Clean up the content of the directory named dist.
    - Post: Create a new blog post using environment variables POST_TITLE and POST_NAME.

To assist with these steps, a Makefile with a series of defined commands is included. Let's review how to use them:

### Build
To build the website just need to run:

```makefile
make build
```

This command will generate the website in the dist/ directory.

### Clean
To clean up the dist/ directory just need to run:

```makefile
make clean
```

This command will remove the contents of the dist/ directory.

### Post
To create a new blog post, set the POST_NAME and POST_TITLE environment variables and run:

```makefile
make post
```

or run the command and name the variables at the same time:

```makefile
make POST_NAME=file_name_example POST_TITLE="Post Title" post
```

This command will create a new blog post with the specified file name and title

### Help
To display a list of available targets and their usage, run:

```makefile
make help
```

This command will show the list of targets and their usage based on the comments in the Makefile.