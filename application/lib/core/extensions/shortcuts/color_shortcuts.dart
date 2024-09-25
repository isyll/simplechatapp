import 'package:flutter/material.dart';

extension ColorsShortcuts on BuildContext {
  ColorScheme get colorScheme => Theme.of(this).colorScheme;
  Brightness get brightness => colorScheme.brightness;
  Color get primaryColor => colorScheme.primary;
  Color get onPrimaryColor => colorScheme.onPrimary;
  Color get secondaryColor => colorScheme.secondary;
  Color get onSecondaryColor => colorScheme.onSecondary;
  Color get successColor => colorScheme.tertiary;
  Color get onSuccessColor => colorScheme.onTertiary;
  Color get errorColor => colorScheme.error;
  Color get onErrorColor => colorScheme.onError;
  Color get surfaceColor => colorScheme.surface;
  Color get onSurfaceColor => colorScheme.onSurface;
  Color get outlineColor => colorScheme.outline;
}
