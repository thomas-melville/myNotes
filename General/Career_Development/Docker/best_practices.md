# best practices

## Incremental build time

1. order matters for caching
  * order your steps from least to most frequently changing
  * copying your jar in goes last.
2. more specific COPY to limit cache busts
  * only copy what's needed.
3. identify cacheable units
  * anything that installs/updates software in the image
  * update the cache and install any packages in one RUN command

## Reduce image size

4. remove unnecessary dependencies
  * apt, use the --no-install-recommends flag
  * only install debugging tools when required.
  * microdnf, don't install the docs
5. remove package manager cache
  * remove it in the same run instruction as the install
  * in apt: rm -rf /var/lib/apt/lists/*

## Maintainability

6. use official images when possible
  * example, instead of installing java use the openjdk image
7. use more specific tags
  * do not use latest
8. look for minimal flavours
  * example, an openjdk image with just the jre, no jdk

## Reproducibility

9. Build from source in a consistent environment
  * build your jar file in the Dockerfile
10. Fetch dependencies in a separate step
  * after copying in the pom resolve the dependencies before bringing in the source code
  * this thinks of cacheability as the pom will change less than the source code.
11. Use multi-stage builds to remove build dependencies
  * build everything in one image
  * then pull only the required artifacts into a new image from the first image
