// import React from 'react';
// import { Tabs } from 'expo-router';
// import { FontAwesome } from '@expo/vector-icons';
// import { AntDesign } from '@expo/vector-icons';

// export default function TabsLayout() {
//   return (
//     <Tabs
//       screenOptions={{
//         tabBarActiveTintColor: 'purple',
//       }}
//     >
//       <Tabs.Screen
//         name="allConnections"
//         options={{
//           title: 'My Connections',
//           tabBarIcon: ({ color }) => <AntDesign size={28} name="contacts" color={color} />,
//         }}
//       />
//       <Tabs.Screen
//         name="index"
//         options={{
//           title: 'Home',
//           tabBarIcon: ({ color }) => <FontAwesome size={28} name="home" color={color} />,
//         }}
//       />
//       <Tabs.Screen
//         name="addConnection"
//         options={{
//           title: 'Add Connection',
//           tabBarIcon: ({ color }) => <AntDesign size={28} name="adduser" color={color} />,
//         }}
//       />
//     </Tabs>
//   );
// }

// app/(tabs)/_layout.tsx
import React from 'react';
import { Tabs } from 'expo-router';
import { FontAwesome } from '@expo/vector-icons';
import { AntDesign } from '@expo/vector-icons';

export default function TabsLayout() {
  return (
    <Tabs
      screenOptions={{
        tabBarActiveTintColor: 'purple',
      }}
    >
      <Tabs.Screen
        name="index" // This will be the "Home" tab
        options={{
          title: 'Home',
          tabBarIcon: ({ color }) => <FontAwesome size={28} name="home" color={color} />,
        }}
      />
      <Tabs.Screen
        name="allConnections"
        options={{
          title: 'My Connections',
          tabBarIcon: ({ color }) => <AntDesign size={28} name="contacts" color={color} />,
        }}
      />
      <Tabs.Screen
        name="addConnection"
        options={{
          title: 'Add Connection',
          tabBarIcon: ({ color }) => <AntDesign size={28} name="adduser" color={color} />,
        }}
      />
    </Tabs>
  );
}
