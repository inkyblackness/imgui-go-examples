## Contributing Guidelines

Thank you for considering to contribute to this project!

The following text lists guidelines for contributions.
These guidelines don't have legal status, so use them as a reference and common sense - and feel free to update them as well!


### "I just want to know..."

For questions, or general usage-information about **Dear ImGui**, please refer to the [homepage](https://github.com/ocornut/imgui), or look in the detailed documentation of the C++ source.
This project contains the "ported" examples of **Dear ImGui**, adopted to a more Go-like style, using the available Go wrappers.


### Extensions

If you can and want to make use of the code from these examples in your own projects, you are happy to do so.

> The code shared between the examples is put into `internal` package on purpose.
> This project contains only reference implementations, not meant to be directly included in others.

Pull-requests with extensions are happily accepted, provided that they uphold the following minimum requirements:
* Code is properly formatted & linted (use [golangci-lint](https://github.com/golangci/golangci-lint) for a full check)

> If there are linter errors that you didn't introduce, you don't have to clean them up - I might have missed them and will be handling them separately.
