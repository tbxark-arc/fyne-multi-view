# Fyne Multi-View Management Examples

This repository contains a collection of examples demonstrating different ways to manage multiple views/pages in a [Fyne](https://fyne.io/) application.

## Prerequisites

- [Go](https://golang.org/doc/install) 1.25+
- [Fyne](https://fyne.io/started/) requirements

## Examples

| ID | Name                 | Description                                                                                                     | Run Command                      |
|:---|:---------------------|:----------------------------------------------------------------------------------------------------------------|:---------------------------------|
| 0  | **AppTabs**          | Standard tabbed interface suitable for multi-document or multi-functional apps.                                 | `go run ./cmd/00_apptabs`        |
| 1  | **DocTabs**          | Similar to AppTabs but with a style more suited for document tabs (like a browser).                             | `go run ./cmd/01_doctabs`        |
| 2  | **Stack**            | Layered views where visibility is manually toggled (Show/Hide). Good for routing.                               | `go run ./cmd/02_stack`          |
| 3  | **Stack (Replace)**  | Uses `container.NewStack` and replaces the `Objects` slice to switch views. Replaces the deprecated Max layout. | `go run ./cmd/03_stack_replace`  |
| 4  | **Border + Sidebar** | Classic desktop layout with a left sidebar menu controlling the main content area.                              | `go run ./cmd/04_border_sidebar` |
| 5  | **Split**            | Resizable split view (e.g., IDE layout).                                                                        | `go run ./cmd/05_split`          |
| 6  | **Accordion**        | Collapsible panels, often used for settings or properties.                                                      | `go run ./cmd/06_accordion`      |
| 7  | **Multi-Window**     | Opening independent secondary windows (e.g., Settings).                                                         | `go run ./cmd/07_multiwindow`    |
| 8  | **Dialog**           | Using modal dialogs for temporary views or confirmations.                                                       | `go run ./cmd/08_dialog`         |
| 9  | **Navigator**        | A custom `ViewManager` struct to encapsulate routing logic. Recommended for larger apps.                        | `go run ./cmd/09_navigator`      |

## Usage

Clone the repository and run the desired example using the commands listed above.

```bash
# Example: Run the Stack demo
go run ./cmd/02_stack
```

## Structure

Each example is located in its own package under the `cmd/` directory.

```
cmd/
  ├── 00_apptabs/
  ├── 01_doctabs/
  ├── ...
  └── 09_navigator/
```
