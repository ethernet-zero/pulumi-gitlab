// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectEnvironmentArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectEnvironmentArgs Empty = new ProjectEnvironmentArgs();

    /**
     * Place to link to for this environment.
     * 
     */
    @Import(name="externalUrl")
    private @Nullable Output<String> externalUrl;

    /**
     * @return Place to link to for this environment.
     * 
     */
    public Optional<Output<String>> externalUrl() {
        return Optional.ofNullable(this.externalUrl);
    }

    /**
     * The name of the environment.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the environment.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID or full path of the project to environment is created for.
     * 
     */
    @Import(name="project", required=true)
    private Output<String> project;

    /**
     * @return The ID or full path of the project to environment is created for.
     * 
     */
    public Output<String> project() {
        return this.project;
    }

    /**
     * Determines whether the environment is attempted to be stopped before the environment is deleted.
     * 
     */
    @Import(name="stopBeforeDestroy")
    private @Nullable Output<Boolean> stopBeforeDestroy;

    /**
     * @return Determines whether the environment is attempted to be stopped before the environment is deleted.
     * 
     */
    public Optional<Output<Boolean>> stopBeforeDestroy() {
        return Optional.ofNullable(this.stopBeforeDestroy);
    }

    private ProjectEnvironmentArgs() {}

    private ProjectEnvironmentArgs(ProjectEnvironmentArgs $) {
        this.externalUrl = $.externalUrl;
        this.name = $.name;
        this.project = $.project;
        this.stopBeforeDestroy = $.stopBeforeDestroy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectEnvironmentArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectEnvironmentArgs $;

        public Builder() {
            $ = new ProjectEnvironmentArgs();
        }

        public Builder(ProjectEnvironmentArgs defaults) {
            $ = new ProjectEnvironmentArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param externalUrl Place to link to for this environment.
         * 
         * @return builder
         * 
         */
        public Builder externalUrl(@Nullable Output<String> externalUrl) {
            $.externalUrl = externalUrl;
            return this;
        }

        /**
         * @param externalUrl Place to link to for this environment.
         * 
         * @return builder
         * 
         */
        public Builder externalUrl(String externalUrl) {
            return externalUrl(Output.of(externalUrl));
        }

        /**
         * @param name The name of the environment.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the environment.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID or full path of the project to environment is created for.
         * 
         * @return builder
         * 
         */
        public Builder project(Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID or full path of the project to environment is created for.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param stopBeforeDestroy Determines whether the environment is attempted to be stopped before the environment is deleted.
         * 
         * @return builder
         * 
         */
        public Builder stopBeforeDestroy(@Nullable Output<Boolean> stopBeforeDestroy) {
            $.stopBeforeDestroy = stopBeforeDestroy;
            return this;
        }

        /**
         * @param stopBeforeDestroy Determines whether the environment is attempted to be stopped before the environment is deleted.
         * 
         * @return builder
         * 
         */
        public Builder stopBeforeDestroy(Boolean stopBeforeDestroy) {
            return stopBeforeDestroy(Output.of(stopBeforeDestroy));
        }

        public ProjectEnvironmentArgs build() {
            $.project = Objects.requireNonNull($.project, "expected parameter 'project' to be non-null");
            return $;
        }
    }

}
