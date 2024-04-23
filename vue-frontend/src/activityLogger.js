export async function logUserActivity(activity,userId) {
    const response = await fetch('http://localhost:8181/userLog', {
        method: 'POST',
        headers: {
            "Content-Type": "application/x-www-form-urlencoded"
        },    
        body: new URLSearchParams({
            'userId': userId,
            'logType': activity,
            'ipAddress': "192.168.1.1"
        })
    })
    if (response.ok) {
        const data = await response.json()
        if (data.status == '200'){
            console.log("Log success!", data.message);
        } else {
            console.error("Failed!", data.message);
        }
    }
}