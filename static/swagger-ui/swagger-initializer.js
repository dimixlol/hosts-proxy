window.onload = function() {
  window.ui = SwaggerUIBundle({
    url: "/api/swagger.json",
    dom_id: '#swagger-ui',
    deepLinking: true,
    presets: [
      SwaggerUIBundle.presets.apis,
    ],
    requestInterceptor: (req) => {
      const cookieName = "__persister_csrf__"
      const header = document.cookie
          .split("; ")
          .find((item) => item.startsWith(cookieName))
      if (header) {
        req.headers["X-CSRF-TOKEN"] = header.split(`${cookieName}=`)[1]
      }
      return req;
    }
  });
};
