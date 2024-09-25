class AppRegex {
  static final namePattern = RegExp(r'[a-zA-Z]+');
  static final emailPattern = RegExp(r'^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$');
}
