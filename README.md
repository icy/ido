## Description

`ido` (`I do`) executes your shell command provided as its input,
but it may wait for you to confirm when there is some potential
risky pattern in your arguments (`live`, `remove`, `delete`,...)

## Build or installation

    $ go get github.com/icy/ido

## Examples

    $ go get github.com/icy/ido
    $ ido rm -rfv /prod/

    :: Found risky pattern: rm
    :: Please type YES and enter to continue:
    :: Thanks, you may have saved your system.

    $ ido echo rm -rfv /prod/

    :: Found risky pattern: rm
    :: Please type YES and enter to continue: YES
    :: Going to execute your command... Best luck.
    rm -rfv /prod/

## Default pattern

    var regExpDefault = regexp.MustCompile("(?i)(del|delete|remove|rm|live|prod|production|format)")

## TODO

- [ ] Allow user to adjust risk pattern
- [ ] Allow external configuration?
- [ ] Print random number as confirmation PIN code

## Authors. License

- [ ] Ky-Anh Huynh
- [ ] `MIT` -- Do almost whatever you want.
