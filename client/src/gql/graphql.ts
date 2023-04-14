/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type Amenity = {
  __typename?: 'Amenity';
  createdAt?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  name: Scalars['String'];
  propertyId: Scalars['ID'];
  provider?: Maybe<Scalars['String']>;
  updatedAt?: Maybe<Scalars['Time']>;
};

export type AmenityInput = {
  name: Scalars['String'];
  propertyId: Scalars['ID'];
  provider: Scalars['String'];
};

export type Bedroom = {
  __typename?: 'Bedroom';
  bedroomNumber: Scalars['Int'];
  createdAt?: Maybe<Scalars['Time']>;
  enSuite: Scalars['Boolean'];
  id: Scalars['ID'];
  master: Scalars['Boolean'];
  propertyUnitId: Scalars['ID'];
  updatedAt?: Maybe<Scalars['Time']>;
};

export type ListingsInput = {
  maxPrice?: InputMaybe<Scalars['Int']>;
  minPrice?: InputMaybe<Scalars['Int']>;
  propertyType?: InputMaybe<Scalars['String']>;
  town: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  addAmenity: Amenity;
  addPropertyUnit: PropertyUnit;
  addPropertyUnitTenant: Tenant;
  addUnitBedrooms: Array<Bedroom>;
  createProperty: Property;
  signIn: Token;
};


export type MutationAddAmenityArgs = {
  input: AmenityInput;
};


export type MutationAddPropertyUnitArgs = {
  input: PropertyUnitInput;
};


export type MutationAddPropertyUnitTenantArgs = {
  input: TenancyInput;
};


export type MutationAddUnitBedroomsArgs = {
  input: Array<UnitBedroomInput>;
};


export type MutationCreatePropertyArgs = {
  input: NewProperty;
};


export type MutationSignInArgs = {
  input: NewUser;
};

export type NewProperty = {
  createdBy: Scalars['ID'];
  maxPrice: Scalars['Int'];
  minPrice: Scalars['Int'];
  name: Scalars['String'];
  postalCode: Scalars['String'];
  town: Scalars['String'];
  type: Scalars['String'];
};

export type NewUser = {
  avatar: Scalars['String'];
  email: Scalars['String'];
  first_name: Scalars['String'];
  last_name: Scalars['String'];
};

export type Property = {
  __typename?: 'Property';
  amenities: Array<Amenity>;
  createdAt?: Maybe<Scalars['Time']>;
  createdBy: Scalars['ID'];
  id: Scalars['ID'];
  maxPrice: Scalars['Int'];
  minPrice: Scalars['Int'];
  name: Scalars['String'];
  owner: User;
  postalCode: Scalars['String'];
  town: Scalars['String'];
  type: Scalars['String'];
  units: Array<PropertyUnit>;
  updatedAt?: Maybe<Scalars['Time']>;
};

export type PropertyUnit = {
  __typename?: 'PropertyUnit';
  bathrooms: Scalars['Int'];
  bedrooms: Array<Bedroom>;
  createdAt?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  propertyId: Scalars['ID'];
  tenancy: Array<Tenant>;
  updatedAt?: Maybe<Scalars['Time']>;
};

export type PropertyUnitInput = {
  bathrooms: Scalars['Int'];
  propertyId: Scalars['ID'];
};

export type Query = {
  __typename?: 'Query';
  getListings: Array<Property>;
  getProperty: Property;
  getUser: User;
  hello: Scalars['String'];
};


export type QueryGetListingsArgs = {
  input: ListingsInput;
};


export type QueryGetPropertyArgs = {
  id: Scalars['ID'];
};


export type QueryGetUserArgs = {
  id: Scalars['ID'];
};

export type TenancyInput = {
  endDate?: InputMaybe<Scalars['Time']>;
  propertyUnitId: Scalars['ID'];
  startDate: Scalars['Time'];
};

export type Tenant = {
  __typename?: 'Tenant';
  createdAt?: Maybe<Scalars['Time']>;
  endDate?: Maybe<Scalars['Time']>;
  id: Scalars['ID'];
  propertyUnitId: Scalars['ID'];
  startDate: Scalars['Time'];
  updatedAt?: Maybe<Scalars['Time']>;
};

export type Token = {
  __typename?: 'Token';
  token: Scalars['String'];
};

export type UnitBedroomInput = {
  bedroomNumber: Scalars['Int'];
  enSuite: Scalars['Boolean'];
  master: Scalars['Boolean'];
  propertyUnitId: Scalars['ID'];
};

export type User = {
  __typename?: 'User';
  avatar: Scalars['String'];
  createdAt?: Maybe<Scalars['Time']>;
  email: Scalars['String'];
  first_name: Scalars['String'];
  id: Scalars['ID'];
  last_name: Scalars['String'];
  properties: Array<Property>;
  updatedAt?: Maybe<Scalars['Time']>;
};

export type HelloQueryVariables = Exact<{ [key: string]: never; }>;


export type HelloQuery = { __typename?: 'Query', hello: string };


export const HelloDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"Hello"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"hello"}}]}}]} as unknown as DocumentNode<HelloQuery, HelloQueryVariables>;
export const namedOperations = {
  Query: {
    Hello: 'Hello'
  }
}