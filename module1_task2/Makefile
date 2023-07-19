.PHONY = build clean post

build:
	hugo

clean:
	rm -rf dist/

post:
	@if [ -z "$(POST_NAME)" ]; then \
		echo "Error: POST_NAME environment variable is not set."; \
		exit 1; \
	fi
	@if [ -z "$(POST_TITLE)" ]; then \
		echo "Error: POST_TITLE environment variable is not set."; \
		exit 1; \
	fi
	hugo new posts/$(POST_NAME).md
	echo "---\ntitle: \"$(POST_TITLE)\"\ndate: `date -Iseconds`\ndraft: true\n---\n" > content/posts/$(POST_NAME).md
	