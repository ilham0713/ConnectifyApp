import React from 'react';
import { SafeAreaView, StyleSheet } from 'react-native';
import CheckInList from './checkInList';

const App = () => {
  // Placeholder data for now
  const fillerData = {
    names: ["Wormy", "Fank", "Eeham"],
    formerCheckins: ["8-2-2024", "3-5-2024", "10-1-2024"],
  };

  return (
    <SafeAreaView style={styles.container}>
      <CheckInList data={fillerData} />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
});

export default App;
