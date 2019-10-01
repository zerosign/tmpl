extern crate criterion;
extern crate lazy_static;
extern crate tmpl;

use combine::{char::string, parser::Parser};
use criterion::{
    black_box, criterion_group, criterion_main, BenchmarkId, Criterion, ParameterizedBenchmark,
    Throughput,
};
use lazy_static::lazy_static;
use std::{borrow::Cow, iter::FromIterator, str};

use tmpl::util::closure;

pub fn closure_simple(c: &mut Criterion) {
    let data : Vec<&'static str> = vec![
        "(((true)))",
        "((((((true))))))",
        "((((((((((((true))))))))))))",
        "((((((((((((((((((((((((true))))))))))))))))))))))))",
        "((((((((((((((((((((((((((((((((((((((((((((((((true))))))))))))))))))))))))))))))))))))))))))))))))",
        "((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((true))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))",
        "((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((true))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))))",
    ];

    c.bench(
        "closure_benchmark",
        ParameterizedBenchmark::new(
            "closure_test",
            |b, param| b.iter(|| closure::<_, _, &str>(string("true"), &('(', ')')).parse(*param)),
            data,
        ),
    );
}

criterion_group!(benches, closure_benchmark);
criterion_main!(benches);
