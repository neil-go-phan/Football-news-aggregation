FROM docker.elastic.co/elasticsearch/elasticsearch:7.12.0
LABEL maintainer="Your Name <yourname@example.com>"

# Install any plugins you need, e.g.
RUN elasticsearch-plugin install analysis-icu

# Set heap size
ENV ES_JAVA_OPTS="-Xms1g -Xmx1g"

# Copy any configuration files you need, e.g.
COPY elasticsearch.yml /usr/share/elasticsearch/config/

# Allow Docker healthcheck to work
HEALTHCHECK --interval=30s --timeout=10s --retries=5 \
  CMD curl -f http://localhost:9200/_cluster/health || exit 1