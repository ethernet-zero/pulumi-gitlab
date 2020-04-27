# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, avatar_url=None, bio=None, can_create_group=None, can_create_project=None, color_scheme_id=None, created_at=None, current_sign_in_at=None, email=None, extern_uid=None, external=None, id=None, is_admin=None, last_sign_in_at=None, linkedin=None, location=None, name=None, organization=None, projects_limit=None, skype=None, state=None, theme_id=None, twitter=None, two_factor_enabled=None, user_id=None, user_provider=None, username=None, website_url=None):
        if avatar_url and not isinstance(avatar_url, str):
            raise TypeError("Expected argument 'avatar_url' to be a str")
        __self__.avatar_url = avatar_url
        """
        The avatar URL of the user.
        """
        if bio and not isinstance(bio, str):
            raise TypeError("Expected argument 'bio' to be a str")
        __self__.bio = bio
        """
        The bio of the user.
        """
        if can_create_group and not isinstance(can_create_group, bool):
            raise TypeError("Expected argument 'can_create_group' to be a bool")
        __self__.can_create_group = can_create_group
        """
        Whether the user can create groups.
        """
        if can_create_project and not isinstance(can_create_project, bool):
            raise TypeError("Expected argument 'can_create_project' to be a bool")
        __self__.can_create_project = can_create_project
        """
        Whether the user can create projects.
        """
        if color_scheme_id and not isinstance(color_scheme_id, float):
            raise TypeError("Expected argument 'color_scheme_id' to be a float")
        __self__.color_scheme_id = color_scheme_id
        """
        User's color scheme ID.
        """
        if created_at and not isinstance(created_at, str):
            raise TypeError("Expected argument 'created_at' to be a str")
        __self__.created_at = created_at
        """
        Date the user was created at.
        """
        if current_sign_in_at and not isinstance(current_sign_in_at, str):
            raise TypeError("Expected argument 'current_sign_in_at' to be a str")
        __self__.current_sign_in_at = current_sign_in_at
        """
        Current user's sign-in date.
        """
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        __self__.email = email
        """
        The e-mail address of the user.
        """
        if extern_uid and not isinstance(extern_uid, str):
            raise TypeError("Expected argument 'extern_uid' to be a str")
        __self__.extern_uid = extern_uid
        """
        The external UID of the user.
        """
        if external and not isinstance(external, bool):
            raise TypeError("Expected argument 'external' to be a bool")
        __self__.external = external
        """
        Whether the user is external.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if is_admin and not isinstance(is_admin, bool):
            raise TypeError("Expected argument 'is_admin' to be a bool")
        __self__.is_admin = is_admin
        """
        Whether the user is an admin.
        """
        if last_sign_in_at and not isinstance(last_sign_in_at, str):
            raise TypeError("Expected argument 'last_sign_in_at' to be a str")
        __self__.last_sign_in_at = last_sign_in_at
        """
        Last user's sign-in date.
        """
        if linkedin and not isinstance(linkedin, str):
            raise TypeError("Expected argument 'linkedin' to be a str")
        __self__.linkedin = linkedin
        """
        Linkedin profile of the user.
        """
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        __self__.location = location
        """
        The location of the user.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        """
        The name of the user.
        """
        if organization and not isinstance(organization, str):
            raise TypeError("Expected argument 'organization' to be a str")
        __self__.organization = organization
        """
        The organization of the user.
        """
        if projects_limit and not isinstance(projects_limit, float):
            raise TypeError("Expected argument 'projects_limit' to be a float")
        __self__.projects_limit = projects_limit
        """
        Number of projects the user can create.
        """
        if skype and not isinstance(skype, str):
            raise TypeError("Expected argument 'skype' to be a str")
        __self__.skype = skype
        """
        Skype username of the user.
        """
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        __self__.state = state
        """
        Whether the user is active or blocked.
        """
        if theme_id and not isinstance(theme_id, float):
            raise TypeError("Expected argument 'theme_id' to be a float")
        __self__.theme_id = theme_id
        """
        User's theme ID.
        """
        if twitter and not isinstance(twitter, str):
            raise TypeError("Expected argument 'twitter' to be a str")
        __self__.twitter = twitter
        """
        Twitter username of the user.
        """
        if two_factor_enabled and not isinstance(two_factor_enabled, bool):
            raise TypeError("Expected argument 'two_factor_enabled' to be a bool")
        __self__.two_factor_enabled = two_factor_enabled
        """
        Whether user's two factor auth is enabled.
        """
        if user_id and not isinstance(user_id, float):
            raise TypeError("Expected argument 'user_id' to be a float")
        __self__.user_id = user_id
        if user_provider and not isinstance(user_provider, str):
            raise TypeError("Expected argument 'user_provider' to be a str")
        __self__.user_provider = user_provider
        """
        The UID provider of the user.
        """
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        __self__.username = username
        """
        The username of the user.
        """
        if website_url and not isinstance(website_url, str):
            raise TypeError("Expected argument 'website_url' to be a str")
        __self__.website_url = website_url
        """
        User's website URL.
        """
class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            avatar_url=self.avatar_url,
            bio=self.bio,
            can_create_group=self.can_create_group,
            can_create_project=self.can_create_project,
            color_scheme_id=self.color_scheme_id,
            created_at=self.created_at,
            current_sign_in_at=self.current_sign_in_at,
            email=self.email,
            extern_uid=self.extern_uid,
            external=self.external,
            id=self.id,
            is_admin=self.is_admin,
            last_sign_in_at=self.last_sign_in_at,
            linkedin=self.linkedin,
            location=self.location,
            name=self.name,
            organization=self.organization,
            projects_limit=self.projects_limit,
            skype=self.skype,
            state=self.state,
            theme_id=self.theme_id,
            twitter=self.twitter,
            two_factor_enabled=self.two_factor_enabled,
            user_id=self.user_id,
            user_provider=self.user_provider,
            username=self.username,
            website_url=self.website_url)

def get_user(email=None,user_id=None,username=None,opts=None):
    """
    Provides details about a specific user in the gitlab provider. Especially the ability to lookup the id for linking to other resources.




    :param str email: The e-mail address of the user. (Requires administrator privileges)
    :param float user_id: The ID of the user.
    :param str username: The username of the user.
    """
    __args__ = dict()


    __args__['email'] = email
    __args__['userId'] = user_id
    __args__['username'] = username
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gitlab:index/getUser:getUser', __args__, opts=opts).value

    return AwaitableGetUserResult(
        avatar_url=__ret__.get('avatarUrl'),
        bio=__ret__.get('bio'),
        can_create_group=__ret__.get('canCreateGroup'),
        can_create_project=__ret__.get('canCreateProject'),
        color_scheme_id=__ret__.get('colorSchemeId'),
        created_at=__ret__.get('createdAt'),
        current_sign_in_at=__ret__.get('currentSignInAt'),
        email=__ret__.get('email'),
        extern_uid=__ret__.get('externUid'),
        external=__ret__.get('external'),
        id=__ret__.get('id'),
        is_admin=__ret__.get('isAdmin'),
        last_sign_in_at=__ret__.get('lastSignInAt'),
        linkedin=__ret__.get('linkedin'),
        location=__ret__.get('location'),
        name=__ret__.get('name'),
        organization=__ret__.get('organization'),
        projects_limit=__ret__.get('projectsLimit'),
        skype=__ret__.get('skype'),
        state=__ret__.get('state'),
        theme_id=__ret__.get('themeId'),
        twitter=__ret__.get('twitter'),
        two_factor_enabled=__ret__.get('twoFactorEnabled'),
        user_id=__ret__.get('userId'),
        user_provider=__ret__.get('userProvider'),
        username=__ret__.get('username'),
        website_url=__ret__.get('websiteUrl'))
