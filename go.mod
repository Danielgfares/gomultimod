module dgf.io/client

go 1.21

replace dgf.io/producer => ./producer

replace dgf.io/consumer => ./consumer

replace dgf.io/pipe => ./pipe

require (
	dgf.io/consumer v0.0.0-00010101000000-000000000000
	dgf.io/pipe v0.0.0-00010101000000-000000000000
	dgf.io/producer v0.0.0-00010101000000-000000000000
)
