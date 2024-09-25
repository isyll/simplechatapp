import 'package:flutter/material.dart';

extension WidgetExtension on State {
  void onWidgetDidBuild(Function callback) {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      callback();
    });
  }
}
