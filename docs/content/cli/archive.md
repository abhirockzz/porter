---
title: "porter archive"
slug: porter_archive
url: /cli/porter_archive/
---
## porter archive

Archive a bundle

### Synopsis

Archives a bundle by generating a gzipped tar archive containing the bundle, invocation image and any referenced images.

```
porter archive [flags]
```

### Examples

```
  porter archive [FILENAME]
  porter archive --file another/porter.yaml [FILENAME]
  porter archive --cnab-file some/bundle.json [FILENAME]
		  
```

### Options

```
      --cnab-file string   Path to the CNAB bundle.json file.
  -f, --file porter.yaml   Path to the Porter manifest. Defaults to porter.yaml in the current directory.
  -h, --help               help for archive
```

### Options inherited from parent commands

```
      --debug   Enable debug logging
```

### SEE ALSO

* [porter](/cli/porter/)	 - I am porter 👩🏽‍✈️, the friendly neighborhood CNAB authoring tool

