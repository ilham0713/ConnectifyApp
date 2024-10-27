import React, { useState } from 'react';
import { View, Text, StyleSheet } from 'react-native';
import Checkbox from 'expo-checkbox';


// Define the props type for the data
type CheckInListProps = {
  data: {
    names: string[];
    formerCheckins: string[];
  };
};

const CheckInList = ({ data }: CheckInListProps) => {
  // Initialize checked state for each item
  const [checkedItems, setCheckedItems] = useState<boolean[]>(new Array(data.names.length).fill(false));

  // Toggle checked state for a specific item
  const toggleCheckbox = (index: number) => {
    setCheckedItems((prev) => {
      const updatedCheckedItems = [...prev];
      updatedCheckedItems[index] = !updatedCheckedItems[index];
      return updatedCheckedItems;
    });
  };

  return (
    <View style={styles.container}>
      {/* Header Row */}
      <View style={styles.headerRow}>
        <Text style={styles.headerText}>Connection</Text>
        <Text style={styles.headerText}>Last Checked</Text>
        <Text style={styles.headerText}>Check-In</Text>
      </View>

      {/* Data Rows */}
      {data.names.map((name: string, index: number) => (
        <View key={index} style={styles.dataRow}>
          <Checkbox
            value={checkedItems[index]}
            onValueChange={() => toggleCheckbox(index)}
            style={styles.checkbox}
          />
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
    textAlign: 'center',
  },
  dataRow: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    paddingVertical: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#eee',
    alignItems: 'center',
  },
  dataText: {
    fontSize: 14,
    flex: 1,
    textAlign: 'center',
  },
  checkbox: {
    marginRight: 10,
  },
});

export default CheckInList;

