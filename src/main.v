module main

import net
import io

fn main() {
	host_addr := 'google.com:80'

	mut conn := net.dial_tcp(host_addr)!
	defer {
		conn.close() or {}
	}

	println('Peer address: ${conn.peer_addr()!}')
	println('Local address: ${conn.addr()!}\r\n\r\n')

	// Send HEAD request for a file
	conn.write_string('HEAD /index.html HTTP/1.0 \r\n\r\n') or {
		eprintln('Failed to send HEAD request to ${host_addr}, error: ${err}')
	}

	// Read all incoming data
	result := io.read_all(reader: conn)!
	println(result.bytestr())
}
