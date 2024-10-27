/* eslint-disable */
import * as Router from 'expo-router';

export * from 'expo-router';

declare module 'expo-router' {
  export namespace ExpoRouter {
    export interface __routes<T extends string = string> extends Record<string, unknown> {
      StaticRoutes: `/` | `/(tabs)` | `/(tabs)/` | `/(tabs)/addConnection` | `/(tabs)/allConnections` | `/(tabs)/checkInList` | `/_sitemap` | `/addConnection` | `/allConnections` | `/checkInList` | `/settings`;
      DynamicRoutes: never;
      DynamicRouteTemplate: never;
    }
  }
}
