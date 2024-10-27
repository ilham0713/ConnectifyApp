import { Text, View, StyleSheet} from "react-native";
import React from "react";
import CheckInList from "./checkInList";
const data = {
  names: ["Wormy", "Fank", "Eeham"],
  formerCheckins: ["8-2-2024", "3-5-2024", "10-1-2024"],
};
export default function Home() {
  return (
    <View style={styles.container}>
      <Text>Today's Check Ins</Text>
      <View style={styles.container}>
      {/* Header Row */}
      <View style={styles.headerRow}>
        <Text style={styles.headerText}>Connection</Text>
        <Text style={styles.headerText}>Last Checked</Text>
      </View>
      <View key={index}>
          <CheckInList 
            name={name} 
            formerCheckin={data.formerCheckins[index]} 
            nextCheckin={data.nextCheckins[index]} 
          />
        </View>
    </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
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

