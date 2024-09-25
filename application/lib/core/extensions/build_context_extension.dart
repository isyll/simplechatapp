import 'package:flutter/material.dart';

extension BuildContextExtension on BuildContext {
  Size get size => MediaQuery.of(this).size;
  double get height => size.height;
  double get width => size.width;
  ScaffoldFeatureController<SnackBar, SnackBarClosedReason> showSnackBar(
          SnackBar snackBar) =>
      ScaffoldMessenger.of(this).showSnackBar(snackBar);
  Future<dynamic> push(Widget newPage) => Navigator.of(this).push(
        MaterialPageRoute(builder: (_) => newPage),
      );
  Future<dynamic> pushReplacement(Widget newPage, {Object? arguments}) =>
      Navigator.of(this).pushReplacement(MaterialPageRoute(
        builder: (context) => newPage,
      ));
  void pop() => Navigator.of(this).pop();
  bool canPop() => Navigator.canPop(this);
  Future<dynamic> pushReplacementNamed(String routeName, {Object? arguments}) =>
      Navigator.of(this).pushReplacementNamed(routeName, arguments: arguments);
  Future<dynamic> pushNamed(String routeName, {Object? arguments}) =>
      Navigator.of(this).pushNamed(routeName, arguments: arguments);
  Locale get locale => Localizations.localeOf(this);
}
