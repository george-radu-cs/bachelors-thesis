class ThemeService {
  ThemeService._privateConstructor();
  // ...
  static final ThemeData _lightBlueTheme = ThemeData(
    appBarTheme: const AppBarTheme(
      backgroundColor: Colors.blueAccent,
      foregroundColor: Colors.white,
    ),
    colorScheme: ColorScheme.fromSeed(
      seedColor: Colors.blueAccent,
      brightness: Brightness.light,
      primary: Colors.blueAccent,
      secondary: Colors.blueAccent.shade100,
    ),
    iconTheme: const IconThemeData(color: Colors.black87),
    listTileTheme: const ListTileThemeData(
      iconColor: Colors.black87,
      contentPadding: EdgeInsets.symmetric(horizontal: 0),
    ),
    useMaterial3: true,
  );
  // ...
  static final ThemeData _darkBlueTheme = ThemeData(
    appBarTheme: AppBarTheme(
      backgroundColor: Colors.blue.shade700,
      foregroundColor: Colors.black,
    ),
    colorScheme: ColorScheme.fromSeed(
      seedColor: Colors.blue.shade700,
      brightness: Brightness.dark,
      primary: Colors.blue.shade700,
      secondary: Colors.blue.shade200,
    ),
    iconTheme: const IconThemeData(color: Colors.white, opacity: .9),
    listTileTheme: const ListTileThemeData(
      textColor: Colors.black87,
      iconColor: Colors.black87,
      contentPadding: EdgeInsets.symmetric(horizontal: 0),
    ),
    useMaterial3: true,
  );
  // ...
}