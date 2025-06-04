![R2D2 Logo](https://r2d2-lang.kinsta.app/assets/images/r2d2-dark.svg)

# R2D2 - The programming Language

**R2D2** is a programming language designed for building modular JavaScript applications using a clear, explicit, and structured syntax.

It is **written in Go** and **compiles to JavaScript**, embracing a module-based architecture where each module can contain variables and functions. These modules are compiled into native JavaScript objects, making it easy to integrate with existing frontend or backend projects.

R2D2 introduces the concept of **pseudo-functions** — a special type of function that can only contain calls to other functions. This enforces composability and encourages building code through reusable, isolated behaviors.

The language is built with simplicity and clarity in mind. It avoids hidden behavior and favors explicit patterns, making it approachable for those who want to write structured, maintainable code that compiles directly to JavaScript.

You can start using **R2D2** by executing the following command on your system.

## For UNIX based systems

```bash
curl -sSL https://raw.githubusercontent.com/ArturC03/r2d2-cli/main/install.sh | bash
```

## For Windows

```bash
Just use WSL, believe me. I'm doing you a favor.
```

To verify the installation, try executing `r2d2`, if it returns something you're good to go!

## Hello World

Let's try to make a hello world program!

1. Create a `.r2d2` file such as `helloworld.r2d2`
2. Put the following content in the file

```js
module HelloWorld {
  export fn main() {
    console.log("Hello World!");
  }
}
```
3. Now you can choose to compile, run or even convert it into a js file by using on of the following commands

```bash
r2d2 build helloworld.r2d2
``` 

```bash
r2d2 run helloworld.r2d2
``` 

```bash
r2d2 js helloworld.r2d2
```

For more information go to the [site](https://r2d2-lang.kinsta.app/docs). Any doughts just make an issue.
