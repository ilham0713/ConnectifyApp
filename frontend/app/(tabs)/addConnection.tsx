import { View, Text, StyleSheet } from 'react-native';

export default function AddConnection() {
  return (
    <View style={styles.container}>
      <Text>AddConnection</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});