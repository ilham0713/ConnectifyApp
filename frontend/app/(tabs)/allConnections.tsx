import { View, Text, StyleSheet} from 'react-native';

export default function AllConnections() {
  return (
    <View style={styles.container}>
      <Text>AllConnections</Text>
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