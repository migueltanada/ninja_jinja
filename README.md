# Ninja Jinja

## About

Ninja Jinja will let you use template for your config, html, xml, and other fles. It will parse your templates based on property files or environment variables.

## Purpose

To learn go(lols)

## Functional Reqirements

Ninja Jinja will be a CLI Tool. It will accept multiple parameters and will output the transformed file through your operating system's stdout. From there you can redirect the outputs into file or other streams based on your operating system.

### Properties files

It should be able to read from the following:

- properties file - doublequotes(") are read differently strongly recommend to ommit them
  - key=value
- yaml - there is a space( ) after the colon(:)
  - key: "value"
- json
  - {}

### Warning and Logging

- Send warnings for parts of the templates that has no property counterpart.
- Logs

### Parameters

- templateFile(required) - file location of the template
- parseType(required) - get values from either environment variables or propertiesFile or both(env/properties), but properties file always first.
- propertiesfile(optional) - if chosen type is propertiesfile, state the file location of props file in this parameter

## Current Issues

- Katamad

- Viper tool(3rd party tool) can only read Keys in lower case.
  - Possible solution: Create an auto property file reader for yaml/properties/json. This will be very tidious especially some Keys can have multiple values.

  ```yaml
    key1: value
    key2:
    - key2.name: value
      key2.test: value
    - key2.name: value
      key2.test: value
    key3:
    - value
    - value
    - value

  ```

- Not yet tested if text/template package is able to get which Keys are defined in the template. This is to be used for the logging part(warnings for keys w/o values). Also for creating a map from environment variables for only the requried values.

## Current Codes

### Limitations

- Options are not yet complete

- Everything is still "todo" lol

- Hardcoded to read yaml

### Samples

html.tmpl

```html
<html>
<head>
<title>{{.property1}}</title>
</head>
<body>
This is Property1: {{.property1}}
This is Property2: {{.property2}}
This is Property3: {{.property3}}
</body>
</html>
```

Properties

```yaml
Property1: "Title"
Property2: Bahay
Property3: Lupa
```

### Usage

```unix
╰─ go run main.go --template html.tmpl --properties Properties

<html>
<head>
<title>Title</title>
</head>
<body>
This is Property1: Title
This is Property2: Bahay
This is Property3: Lupa
</body>
</html>
```
