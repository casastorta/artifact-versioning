# Artifact Versioning

## Usage

Want to skip all this small talk and read directly how to use this? [Go here](usage.md).

Want to contribute? Check out [developer notes](devnotes.md)<sup>TODO</sup>.

## tl; dr

Artifact Versioning is a [calendar versioning utility](https://calver.org) for your local and CI builds, 
taking in account SCM branching ( `master` / `main` branches are calculated without suffixes).

Currently supported versioning schema is `YYYY.0W[. BRANCH_SUFFIX]. MICRO` , as this is, specifically, 
versioning schema I've needed for my projects. If you need other versioning schema, either request
it via issues or submit a PR and it will be considered for implementation.

## Why?

Based on the situation in real life with projects I work on, which have common versioning requirement
but use different toolkit and different implementations to achieve the same funcitonality. I am 
encountering on almost daily basis versioning calculation written in languages like Groovy (for Gradle
and/or Jenkins), Python, Ruby, different shell flavors and who knows what else is somewhere in
the source code repositories...

I have opted to write an implementation in Go, hoping I could use it as a replacement for all of the
different implementations I run into.

### Why not simply staying with shell scripts?

While reasonably widespread shell (not unlike `bash` ) is relatively universal approach and seems
reasonable for a task like this - I've still felt a need to rewrite it in something else which can 
be used equally or more universally across use-cases and platforms.

First problem with shell scripting is - a platform. While I think I wasn't made to work with any kind
of Windows-based servers for two decades by now (Exchange which I've used via Outlook excluded), and
my personal workstations are not running any flavor of Windows for almost the same time period - I
acknowledge Windows exist. And not everyone is working solely under 
[Windows' WSL](https://docs.microsoft.com/en-us/windows/wsl/about).

Second big problem with shell scripting, when used to implement even remotely complex logic, is 
maintainability. Shell scripting is by now more than half-century old technology, and it shows. 
Also, it is not a full blown programming language by definition, so a lot of relatively routine
operations in any other language are with shell scripting done by invoking external utilities.
That and more results in the fact that most of the developers I've worked with in my life are
not fans of maintaining complex shell scripts (to put it mildly).

### Why Go?

Well, to to be honest - I've mainly wanted to write some Go code as a part of my learning process to
learn new programming language. 

But beyond that, Go has couple of advantages which make it (or its executable outputs) ideal for this
specific use-case and problem I've wanted to solve.

Obviously, even vaguely complex codebase is more maintainable in the "proper" programming language
than when implemented in shell scripts.

Furthermore, Go is truly portable. Many languages today promise 
"[write once, run everywhere](https://en.wikipedia.org/wiki/Write_once,_run_anywhere)", and they
do deliver to some extent. But Go goes beyond that. Foundationally, from the inception, Go supports 
cross-platform compiling and it has been
[only improving ever since](https://dave.cheney.net/2015/03/03/cross-compilation-just-got-a-whole-lot-better-in-go-1-5).
This makes compiling and distribution of the executables built with Go a breeze.
