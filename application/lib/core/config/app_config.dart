import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:simple_chat_app/core/constants/app_keys.dart';

class AppConfig {
  static const appName = 'Simple Chat App';
  static const supportedLocales = [Locale('fr', 'SN')];
  static const defaultLocale = Locale('fr', 'SN');
  static const author = 'Ibrahima Sylla';
  static const authorEmail = 'isyll711@gmail.com';

  static Future<void> init() async {
    final prefs = await SharedPreferences.getInstance();

    if (prefs.get(AppKeys.settings) == null) {
      prefs.setString(AppKeys.settings, '{}');
    }
  }
}
