terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0"
    }
  }
}

provider "docker" {}

resource "docker_image" "nginx" {
  name = "docker.io/library/nginx:alpine"
}

resource "docker_container" "web" {
  name  = "infrapilot-web"
  image = docker_image.nginx.image_id
  ports {
    internal = 80
    external = 8081
  }
}
