# Structure of past project
```
Project
    |_ web
        |_docker
            |_dev
                |_.dev
                |_Dockerfile
                |_requirements.txt
                |_dbenv
                |_ ...
            |_staging
                |_.dev
                |_Dockerfile
                |_requirements.txt
                |_dbenv
                |_ ...
            |_prod
                |_.dev
                |_Dockerfile
                |_requirements.txt
                |_dbenv
                |_ ...
        |_settings
            |___init__.py
            |_base.py
            |_dev.py
            |_staging.py
            |_prod.py
        |_base
            |_migrations
            |___init__.py
            |_admin.py
            |_apps.py
            |_models.py
            |_serializers.py
            |_tests.py
            |_urls.py
            |_views.py
            |_datetime.py
            |_exceptions.py
        |_static
            |_assets
                |_css
                |_fonts
                |_images
                |_js
                |_
            |_templates
                |_base
                    |_base.html
                    |_footer.html
                    |_index.html
                    |_topnav.html
                |_app(s)
        |_manage.py
    |_.gitignore
    |_docker-compose.yaml
    |_docker-compose-staging.yaml
    |_docker-compose-prod.yaml
```

web                 - This is the main directory of all apps in project
docker              - This is the directory that holds all configurations of all environments when building and deploying all containers
                      In docker/staging and docker/staging will be a env file for particular envireonments on servers
settings            - This is where all the settings of Django including all environments(dev, staging, production)
base                - This will be the 1st app created when initializing a new project, this app holds the base classes for generic views, 
                      serializers, exception handler, and utilities   functions like datetime,...
static              - This is where all static files is storing, html, js, css, images, fonts, etc.
static/staticfiles  - This is only existing in staging and production environments when I used gunicorn to serve django service, 
                      this directory will be a root directory for   serving static files

docker-compose.yaml / docker-compose-staging.yaml / docker-compose-prod.yaml - These three compose files are for building containers out of a project which getting configurations from associated environments in docker/... directory


# For defined conventions
1. Static variable will be named all uppercase separated by underscore ex. COUNT, RATE, DEFAULT_VALUE
2. All other variables will be camelCase
3. For function name, the 1st word will be verb describing what is the action of function, and 2nd word will be what is affected. Also camelCase.
4. Any configurations value, mostly static value, will be separated in another file such as config.py
