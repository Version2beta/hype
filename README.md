# Hype

Brief introduction of the project

### Description of what the project does

## Getting Started

### Prerequisites (e.g., Go, Git, Node.js)

### Installation instructions

### Usage instructions (e.g., how to run Hype, how to create new content, how to deploy the site)

### How it works

#### Create

#### Build

#### React

#### Components

Hype supports three kinds of components.

**Partial Components**: Parts of a page that are reused across multiple pages and templates. These include:

- **Header**: a navigation bar at the top of the page
- **Footer**: a footer at the bottom of the page
- **Sidebar**: a column to the side of the main content area
- **Nav**: a navigation menu that may be located in the header or sidebar

**Page Components**: These render the pages themselves. These may not be separate pages in a file system, but they are designed to act as if they are separate pages.

- **Home**: the home page of the site
- **Page**: a generic content page
- **Gallery**: a page that displays a collection of images
- **Post**: a page that displays a single blog post
- **List**: a page that displays a list of blog post titles and summaries
- **Feed**: a page that displays multiple blog posts sequentially, often with infinite scroll
- **ContactForm**: a form that allows users to contact the site owners or administrators
- **SiteMap**: a page that shows the structure of the website, with links to each page

The third kind of components are custom, designed and built by the user or programmer to best present the content for the site. These may be page components.

#### Markdown source code

Headmatter

- Title
- Slug
- Description
- Status (none / draft / published)
- Date
- Tags
- Authors
- Featured image

Markdown follows, describing the content.

An example Markdown file may look like this:

```
---
title: Example
slug: example
description: An example post
featured_image:
  url: https://example.com/image.jpg
  caption: An example image
  credit: John Doe
author: John Smith
tags:
  - example
date: 2022-02-24
---

This is the content of the post.
```

#### Generated JSON files

Example object rendered by Hype at build time, and delivered to the client when requested:

```
{
  "title": "Example Post",
  "slug": "example-post",
  "description": "An example post about Hype",
  "published": true,
  "date": "2022-03-01T13:00:00Z",
  "tags": ["example", "hype"],
  "authors": [
    {
      "name": "John Doe",
      "email": "john@example.com"
    }
  ],
  "featured_image": {
    "src": "/images/example.jpg",
    "alt": "An example image",
    "caption": "This is an example image",
    "credit": "John Doe"
  },
  "components": [
    {
      "type": "Post",
      "props": {
        "content": [
          {
            "type": "heading",
            "content": "Hype is awesome!",
            "props": {
              "level": 1
            }
          },
          {
            "type": "paragraph",
            "content": "Hype is a great tool for building websites quickly and easily."
          },
          {
            "type": "image",
            "props": {
              "src": "/images/hype.png",
              "alt": "The Hype logo",
              "caption": "The Hype logo"
            }
          },
          {
            "type": "list",
            "content": [
              {
                "type": "list-item",
                "content": "Hype is easy to use"
              },
              {
                "type": "list-item",
                "content": "Hype is fast"
              },
              {
                "type": "list-item",
                "content": "Hype is flexible"
              }
            ]
          }
        ]
      }
    }
  ]
}
```

A description of each key:

- "title": The title of the post.
- "slug": A URL-friendly version of the title.
- "description": A brief description of the post's content.
- "published": A boolean indicating whether the post should be displayed on the site.
- "date": The date and time the post was published.
- "tags": A list of tags associated with the post.
- "authors": A list of authors who contributed to the post, including their name and email address.
- "featured_image": An object representing the featured image for the post, including its source, alt text, caption, and credit information.
- "components": An array of components that make up the post, in the order they should be rendered. Each component has a type and props. The type specifies the type of React component to use for this element, and the props are the properties that should be passed to that component. The props may contain a content key which is an array of nested components. In the example above, the components include a heading, a paragraph, an image, and a list.

## Using Hype

### Description of the file structure of the project

```
|- content/
|  |- posts/
|  |  |- my-first-post.md
|  |  |- my-second-post.md
|  |- pages/
|  |  |- about.md
|  |  |- contact.md
|- public/
|  |- index.html
|  |- assets/
|  |  |- css/
|  |  |  |- main.css
|  |  |- js/
|  |  |  |- main.js
|  |  |- images/
|  |  |  |- logo.png
|- src/
|  |- components/
|  |  |- Header.js
|  |  |- Footer.js
|  |  |- Post.js
|  |  |- List.js
|  |  |- Feed.js
|  |- pages/
|  |  |- About.js
|  |  |- Contact.js
|  |- utils/
|  |  |- markdown.js
|  |  |- fetchContent.js
|  |- index.js
|- package.json
|- README.md
|- .gitignore
|- build.js
|- deploy.js
```

- `content/`: This directory contains all the markdown files for the site, organized by content type.
- `public/`: This directory contains all the static assets for the site, such as images, stylesheets, and JavaScript files.
- `src/`: This directory contains all the React components, organized by content type and other functionality.
- `components/`: This directory contains reusable components that are used across multiple pages, such as Header and Footer.
- `pages/`: This directory contains components that represent individual pages on the site, such as About and Contact.
- `utils/`: This directory contains utility functions for the site, such as the markdown function for parsing markdown into JSX and fetchContent for fetching content from the server.
- `package.json`: This file contains information about the project and its dependencies.
- `README.md`: This file contains information about the project, including how to set it up and use it.
- `.gitignore`: This file specifies which files and directories should be ignored by Git when committing changes to the repository.
- `build.js`: This file contains the script for building the site, which parses the markdown files into JSON objects and generates the static site.
- `deploy.js`: This file contains the script for deploying the site to a server or hosting platform.

### Explanation of what each file or directory does

## Configuration

### Description of the various configuration options and how to use them

### Example configuration files

## Supported Features

### List of features supported by Hype (e.g., markdown to HTML, React support)

### Description of each feature and how to use it

## Contributing

### Guidelines for contributing to the project

### Code of Conduct

## License

### Description of the project's license

### Link to the full license text

## Acknowledgements

### List of contributors or organizations that have contributed to the project

### Any relevant acknowledgements or thanks
