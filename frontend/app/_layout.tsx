// import React from 'react';
// import { GestureHandlerRootView } from 'react-native-gesture-handler';
// import { Drawer } from 'expo-router/drawer';
// import { View, Text, TouchableOpacity } from 'react-native';
// import { FontAwesome } from '@expo/vector-icons';
// import { DrawerNavigationProp } from '@react-navigation/drawer';
// import { ParamListBase } from '@react-navigation/native';
// import { SafeAreaView } from 'react-native-safe-area-context';

// // Define the type for the props expected by CustomHeader
// type HeaderProps = {
//   navigation: DrawerNavigationProp<ParamListBase>;
// };

// const CustomHeader: React.FC<HeaderProps> = ({ navigation }) => {
//   return (
//     <SafeAreaView style={{ backgroundColor: '#fff', elevation: 4 }}>
//       <View
//         style={{
//           flexDirection: 'row',
//           justifyContent: 'space-between',
//           alignItems: 'center',
//           padding: 15,
//         }}
//       >
//         <Text style={{ fontSize: 20, fontWeight: 'bold' }}>App Title</Text>
//         {/* Hamburger Icon */}
//         <TouchableOpacity onPress={() => navigation.openDrawer()}>
//           <FontAwesome name="bars" size={28} color="black" />
//         </TouchableOpacity>
//       </View>
//     </SafeAreaView>
//   );
// };

// export default function RootLayout() {
//   return (
//     <GestureHandlerRootView style={{ flex: 1 }}>
//       <Drawer
//         screenOptions={({ navigation }) => ({
//           header: () => <CustomHeader navigation={navigation} />,
//           drawerPosition: 'right', // Set the drawer to open from the right side
//         })}
//       >
//         {/* Define Drawer Screens */}
//         <Drawer.Screen
//           name="tabs"
//           options={{
//             drawerLabel: 'Home',
//           }}
//         />
//         <Drawer.Screen
//           name="settings"
//           options={{
//             drawerLabel: 'Settings',
//           }}
//         />
//       </Drawer>
//     </GestureHandlerRootView>
//   );
// }

// app/_layout.tsx
import React from 'react';
import { GestureHandlerRootView } from 'react-native-gesture-handler';
import { Drawer } from 'expo-router/drawer';
import { View, Text, TouchableOpacity } from 'react-native';
import { FontAwesome } from '@expo/vector-icons';
import { SafeAreaView } from 'react-native-safe-area-context';

// Custom header to open the drawer
const CustomHeader: React.FC<{ navigation: any }> = ({ navigation }) => {
  return (
    <SafeAreaView style={{ backgroundColor: '#fff', elevation: 4 }}>
      <View
        style={{
          flexDirection: 'row',
          justifyContent: 'space-between',
          alignItems: 'center',
          padding: 15,
        }}
      >
        <Text style={{ fontSize: 20, fontWeight: 'bold' }}>Connectify</Text>
        <TouchableOpacity onPress={() => navigation.openDrawer()}>
          <FontAwesome name="bars" size={28} color="black" />
        </TouchableOpacity>
      </View>
    </SafeAreaView>
  );
};

export default function RootLayout() {
  return (
    <GestureHandlerRootView style={{ flex: 1 }}>
      <Drawer
        screenOptions={({ navigation }) => ({
          header: () => <CustomHeader navigation={navigation} />,
          drawerPosition: 'right', // Set the drawer to open from the right side
        })}
      >
        {/* Reference the folder "(tabs)" which contains the tab navigator */}
        <Drawer.Screen
          name="(tabs)" // Reference the folder "(tabs)" for the tab layout
          options={{
            drawerLabel: 'Home',
          }}
        />
        <Drawer.Screen
          name="settings" // Direct reference to the settings screen
          options={{
            drawerLabel: 'Settings',
          }}
        />
      </Drawer>
    </GestureHandlerRootView>
  );
}
