/*
Package core
envset needs rules for interacting with
shells config files, it can be done directly by checking which
shell we are using but that will limit new shell adding, as the
programs grows we need to add more functions and more logic and checking
to work with so it would be better if we can use shell config with simple
rule declaration.

Now why use rules inside of go program.

The reason is simple to avoid security checks
This also can be break but not easily
Most of the software use their own config files
some of are not meant to use by user but still if someone changes those
config, they can effect the system.
for example if they change the path and make it traverse through '../' to go
their desire location and effect and file. or Just copy you env
That is the reason to do this.

Note: When we discover any better way we will implement that
*/
package core

// Rules for working shells
