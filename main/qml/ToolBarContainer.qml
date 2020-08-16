import QtQuick 2.2
import QtQuick.Controls 2.3
import QtQuick.Layouts 1.3

GroupBox {
    title: qsTr("Synchronize")
    label.x: width/2 - label.contentWidth/2
    label.y: height - 10
    anchors.bottomMargin: 10

    default property alias content: gridLayout.children

    GridLayout {
        id: gridLayout
        anchors.fill: parent
        CheckBox { text: qsTr("E-mail") }
        CheckBox { text: qsTr("Calendar") }
        CheckBox { text: qsTr("Contacts") }
    }
}