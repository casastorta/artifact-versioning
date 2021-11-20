# Artifact Versioning Usage

## Getting Help

At any time you can get help for the utility by using either switch `-h` or `--help` when running
it:

```
# artifact-versioning --help
Usage of artifact-versioning:
  -b string
    	Branching detections mechanism to use (default "jenkins")
  -f	Use fallback mechanisms to try and detect branch name (default is false [not set])
  -fm string
    	Fallback mechanisms to be used for branching detection if the main one fails
  -l string
    	Timezone/location info (default "UTC")
  -m int
    	Micro version number (i. e. build number)
  -t string
    	Time you want to use in YYYY-MM-DD format (default will be current time) (default "2021-11-20")
```

## SCM Branch Name Detection Methods

The following methods of detecting SCM branch of the codebase are possible:

1. `git`
2. `bamboo`
3. `jenkins`

Any or all of the methods can be used either as primary or fallback method(s) for detection.

For example, issuing command like this:

```
artifact-versioning -b git
```

...will attempt to detect SCM branch only via `git` mechanism, while issuing command like this:

```
artifact-versioning -b jenkins -f -fm bamboo,git
```

...will attempt to detect SCM branch in the following order: `jenkins`, then `bamboo` and then `git`.

Using parameter `-f` will allow for fallback mechanisms to be used, while continued to the parameter 
`-fm` we are listing methods which should be used as fallback methods, and in the order in which they
should be attempted.

## Micro Version Number

Pass any integer to this parameter, if you need it. IF you leave it out, it will default to `.0`.
