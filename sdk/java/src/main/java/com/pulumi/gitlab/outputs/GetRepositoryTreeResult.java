// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gitlab.outputs.GetRepositoryTreeTree;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRepositoryTreeResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The path inside repository. Used to get content of subdirectories.
     * 
     */
    private @Nullable String path;
    /**
     * @return The ID or full path of the project owned by the authenticated user.
     * 
     */
    private String project;
    /**
     * @return Boolean value used to get a recursive tree (false by default).
     * 
     */
    private @Nullable Boolean recursive;
    /**
     * @return The name of a repository branch or tag.
     * 
     */
    private String ref;
    /**
     * @return The list of files/directories returned by the search
     * 
     */
    private List<GetRepositoryTreeTree> trees;

    private GetRepositoryTreeResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The path inside repository. Used to get content of subdirectories.
     * 
     */
    public Optional<String> path() {
        return Optional.ofNullable(this.path);
    }
    /**
     * @return The ID or full path of the project owned by the authenticated user.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return Boolean value used to get a recursive tree (false by default).
     * 
     */
    public Optional<Boolean> recursive() {
        return Optional.ofNullable(this.recursive);
    }
    /**
     * @return The name of a repository branch or tag.
     * 
     */
    public String ref() {
        return this.ref;
    }
    /**
     * @return The list of files/directories returned by the search
     * 
     */
    public List<GetRepositoryTreeTree> trees() {
        return this.trees;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRepositoryTreeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String path;
        private String project;
        private @Nullable Boolean recursive;
        private String ref;
        private List<GetRepositoryTreeTree> trees;
        public Builder() {}
        public Builder(GetRepositoryTreeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.path = defaults.path;
    	      this.project = defaults.project;
    	      this.recursive = defaults.recursive;
    	      this.ref = defaults.ref;
    	      this.trees = defaults.trees;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder path(@Nullable String path) {
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        @CustomType.Setter
        public Builder recursive(@Nullable Boolean recursive) {
            this.recursive = recursive;
            return this;
        }
        @CustomType.Setter
        public Builder ref(String ref) {
            this.ref = Objects.requireNonNull(ref);
            return this;
        }
        @CustomType.Setter
        public Builder trees(List<GetRepositoryTreeTree> trees) {
            this.trees = Objects.requireNonNull(trees);
            return this;
        }
        public Builder trees(GetRepositoryTreeTree... trees) {
            return trees(List.of(trees));
        }
        public GetRepositoryTreeResult build() {
            final var o = new GetRepositoryTreeResult();
            o.id = id;
            o.path = path;
            o.project = project;
            o.recursive = recursive;
            o.ref = ref;
            o.trees = trees;
            return o;
        }
    }
}
