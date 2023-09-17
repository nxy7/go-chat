#!/usr/bin/env nu

# Returns PWD from root directory
export def project-root [] {
  let ls = (ls -a | where name == ".git");
  if (($ls | length) == 0) {
    (cd ..);
    let current_path = (pwd);
    if ($current_path == "/") {
      error make {msg: "Got to filesystem root without encountering .git"};
    }
    return (project-root);
  } else {
    return (pwd);
  }
}

export def telepresence-join [] {
  telepresence helm install
  telepresence connect
}

export def telepresence-intercept [] {
  telepresence intercept frontend --port 3000:http --mechanism tcp --namespace chat-app
  telepresence intercept backend --port 8282:http --mechanism tcp --namespace chat-app
}
