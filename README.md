# misfits

This is a golang application that uses ebiten for the graphics engine.  The application is split across the following folder structure

```plaintext
.
├── cmd
│   ├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   ├── config.go
│   │   ├── defaultdependencies.go
│   │   ├── defaultdependencies_test.go
│   │   └── globaloptions.go
│   ├── dependencies
│   │   ├── dependencymanager.go
│   │   └── dependencymanager_test.go
│   ├── embeds
│   │   ├── assetmanager.go
│   │   ├── assetmanager_test.go
│   │   └── assets
│   │       ├── fonts
│   │       │   └── DejaVuSans.ttf
│   │       └── images
│   ├── game
│   │   ├── game.go
│   │   ├── game_manager.go
│   │   ├── game_manager_options.go
│   │   ├── game_manager_test.go
│   │   ├── game_mock.go
│   │   └── game_options.go
│   ├── resources
│   │   ├── fonts.go
│   │   └── fonts_test.go
│   └── screenassets
│       ├── screenasset.go
│       ├── states
│       │   ├── menu.go
│       │   └── states.go
│       └── ui
│           ├── button.go
│           └── button_test.go
├── LICENSE
└── README.md
```

## Dependencies

The Dependency Manager should be able to inject the dependencies into any object that is getting created.


## Testing

For the purpose of testing, this application should use the main asset manager.