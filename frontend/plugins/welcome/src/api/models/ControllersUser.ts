/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersUser
 */
export interface ControllersUser {
    /**
     * 
     * @type {string}
     * @memberof ControllersUser
     */
    email?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersUser
     */
    password?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersUser
     */
    userstatus?: number;
}

export function ControllersUserFromJSON(json: any): ControllersUser {
    return ControllersUserFromJSONTyped(json, false);
}

export function ControllersUserFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersUser {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': !exists(json, 'email') ? undefined : json['email'],
        'password': !exists(json, 'password') ? undefined : json['password'],
        'userstatus': !exists(json, 'userstatus') ? undefined : json['userstatus'],
    };
}

export function ControllersUserToJSON(value?: ControllersUser | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'email': value.email,
        'password': value.password,
        'userstatus': value.userstatus,
    };
}

