# Google Wire Example

Google Wire uses a code generation approach, that is it does compile time
dependency injection using `go generate`.

Two core notions:

* providers
* injectors

Providers produce a value, using plain Go code (e.g. a function). They can
specify dependencies (e.g. parameters) and can also return errors.

A number of providers (and sets) can be grouped into a set:

* [ProviderSet](https://pkg.go.dev/github.com/google/wire#ProviderSet)

At construction time, we want to create required values (in a suitable order) -
this is done with an injector. The signature guides the process of creating the
dependent values.

By default, the generated file is named:

* `wire_gen.go`

And it is guarded by a build constraint:

```
//+build !wireinject
```

## Basic Example

* [basic](basic)

## Interfaces

Go takes an implementation-first approach:

> Do not define interfaces before they are used: without a realistic example of
> usage, it is too difficult to see whether an interface is even necessary, let
> alone what methods it ought to contain.

Also:

* DI can be used to match an implementation to an interface
* do not create a constructor, that returns an interface - even though wire uses standard Go type identity rules for matching.
* interface binding should be put into the provider set

* [iface](iface)

## Struct Fields

Inject struct fields.

* [structs](structs)

