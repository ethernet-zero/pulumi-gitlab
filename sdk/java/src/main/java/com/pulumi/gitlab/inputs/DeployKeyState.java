// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gitlab.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DeployKeyState extends com.pulumi.resources.ResourceArgs {

    public static final DeployKeyState Empty = new DeployKeyState();

    /**
     * Allow this deploy key to be used to push changes to the project. Defaults to `false`.
     * 
     */
    @Import(name="canPush")
    private @Nullable Output<Boolean> canPush;

    /**
     * @return Allow this deploy key to be used to push changes to the project. Defaults to `false`.
     * 
     */
    public Optional<Output<Boolean>> canPush() {
        return Optional.ofNullable(this.canPush);
    }

    /**
     * The public ssh key body.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return The public ssh key body.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * The name or id of the project to add the deploy key to.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The name or id of the project to add the deploy key to.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * A title to describe the deploy key with.
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return A title to describe the deploy key with.
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    private DeployKeyState() {}

    private DeployKeyState(DeployKeyState $) {
        this.canPush = $.canPush;
        this.key = $.key;
        this.project = $.project;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeployKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeployKeyState $;

        public Builder() {
            $ = new DeployKeyState();
        }

        public Builder(DeployKeyState defaults) {
            $ = new DeployKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param canPush Allow this deploy key to be used to push changes to the project. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder canPush(@Nullable Output<Boolean> canPush) {
            $.canPush = canPush;
            return this;
        }

        /**
         * @param canPush Allow this deploy key to be used to push changes to the project. Defaults to `false`.
         * 
         * @return builder
         * 
         */
        public Builder canPush(Boolean canPush) {
            return canPush(Output.of(canPush));
        }

        /**
         * @param key The public ssh key body.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key The public ssh key body.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param project The name or id of the project to add the deploy key to.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The name or id of the project to add the deploy key to.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param title A title to describe the deploy key with.
         * 
         * @return builder
         * 
         */
        public Builder title(@Nullable Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title A title to describe the deploy key with.
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public DeployKeyState build() {
            return $;
        }
    }

}
