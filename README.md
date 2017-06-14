racoon is a utility that loops through a specific munki manifest at an interval and install everything available in that manifest.

By default, `racoon` will check the munki manifest `dep_bootstrap` every two seconds(unless the previous run is still going).
The manifest and interval between checks is configurable.

# Usage

```
Usage of ./racoon:
  -interval duration
    	interval between checks (default 2s)
  -manifest string
    	manifest to loop through (default "dep_bootstrap")
```
