import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

const data = {
  names: ["Wormy", "Fank", "Eeham"],
  formerCheckins: ["8-2-2024", "3-5-2024", "10-1-2024"],
  nextCheckins: ["10-24-2024", "10-24-2024", "10-24-2024"],
};

const CheckInList = () => {
  return (
    <View style={styles.container}>
      {/* Header Row */}
      <View style={styles.headerRow}>
        <Text style={styles.headerText}>Connection</Text>
        <Text style={styles.headerText}>Last Checked</Text>
      </View>

      {/* Data Rows */}
      {data.names.map((name: string, index: number) => (
        <View key={index} style={styles.dataRow}>
          <Text style={styles.dataText}>{name}</Text>
          <Text style={styles.dataText}>{data.formerCheckins[index]}</Text>
        </View>
      ))}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    padding: 20,
  },
  headerRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    paddingBottom: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  headerText: {
    fontWeight: 'bold',
    fontSize: 16,
    flex: 1,
  },
  dataRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    paddingVertical: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#eee',
  },
  dataText: {
    fontSize: 14,
    flex: 1,
  },
});

export default CheckInList;
