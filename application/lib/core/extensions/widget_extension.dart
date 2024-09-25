import 'package:flutter/material.dart';

extension WidgetExtension on Widget {
  void onWidgetDidBuild(Function callback) {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      callback();
    });
  }
}
